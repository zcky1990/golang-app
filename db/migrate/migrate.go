// Package migrate mirrors Rails' db/migrate directory.
//
// Every file in this package named <timestamp>_<name>.go registers one
// migration from its init() function. Applied versions are tracked in the
// schema_migrations table, exactly like Rails does with ActiveRecord, so
// `go run ./db/migrate/cmd -direction=up` only runs what is still pending.
package migrate

import (
	"context"
	"fmt"
	"log"
	"sort"
	"sync"

	"gorm.io/gorm"
)

// Migration is a single versioned database change.
type Migration struct {
	// Version is a Rails-style timestamp, e.g. "20240501120000".
	Version string
	// Name is the human readable part of the filename, e.g. "create_users_table".
	Name string
	Up   func(tx *gorm.DB) error
	Down func(tx *gorm.DB) error
}

var (
	mu         sync.Mutex
	migrations []Migration
)

// Register adds a migration to the global list. It is meant to be called from
// an init() in this package so a new file is all you need to add a change.
func Register(migration Migration) {
	mu.Lock()
	defer mu.Unlock()
	for _, existing := range migrations {
		if existing.Version == migration.Version {
			panic(fmt.Sprintf("migration version %s registered twice", migration.Version))
		}
	}
	migrations = append(migrations, migration)
}

// Migrations returns every registered migration ordered by version.
func Migrations() []Migration {
	mu.Lock()
	defer mu.Unlock()
	ordered := make([]Migration, len(migrations))
	copy(ordered, migrations)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Version < ordered[j].Version })
	return ordered
}

// Runner executes migrations against a PostgreSQL database.
type Runner struct {
	db *gorm.DB
}

// NewRunner builds a runner bound to a GORM connection.
func NewRunner(db *gorm.DB) *Runner {
	return &Runner{db: db}
}

const schemaMigrationsTable = "schema_migrations"

// ensureSchemaMigrations creates the version ledger.
func (r *Runner) ensureSchemaMigrations(ctx context.Context) error {
	return r.db.WithContext(ctx).Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version varchar(40) PRIMARY KEY,
			applied_at timestamptz NOT NULL DEFAULT now()
		)`, schemaMigrationsTable)).Error
}

// appliedVersions returns the versions already recorded in the ledger.
func (r *Runner) appliedVersions(ctx context.Context) (map[string]bool, error) {
	var versions []string
	err := r.db.WithContext(ctx).
		Table(schemaMigrationsTable).
		Order("version ASC").
		Pluck("version", &versions).Error
	if err != nil {
		return nil, err
	}
	applied := make(map[string]bool, len(versions))
	for _, version := range versions {
		applied[version] = true
	}
	return applied, nil
}

// Up applies every pending migration in version order.
func (r *Runner) Up(ctx context.Context) error {
	if err := r.ensureSchemaMigrations(ctx); err != nil {
		return fmt.Errorf("could not prepare %s: %w", schemaMigrationsTable, err)
	}
	applied, err := r.appliedVersions(ctx)
	if err != nil {
		return err
	}

	for _, migration := range Migrations() {
		if applied[migration.Version] {
			continue
		}
		log.Printf("[migrate] up   %s_%s", migration.Version, migration.Name)
		if err := r.run(ctx, migration, migration.Up); err != nil {
			return fmt.Errorf("migration %s_%s failed: %w", migration.Version, migration.Name, err)
		}
		if err := r.record(ctx, migration.Version); err != nil {
			return err
		}
	}
	return nil
}

// Down rolls back the most recently applied migration, Rails' `db:rollback`.
func (r *Runner) Down(ctx context.Context) error {
	if err := r.ensureSchemaMigrations(ctx); err != nil {
		return fmt.Errorf("could not prepare %s: %w", schemaMigrationsTable, err)
	}
	applied, err := r.appliedVersions(ctx)
	if err != nil {
		return err
	}

	ordered := Migrations()
	for i := len(ordered) - 1; i >= 0; i-- {
		migration := ordered[i]
		if !applied[migration.Version] {
			continue
		}
		log.Printf("[migrate] down %s_%s", migration.Version, migration.Name)
		if err := r.run(ctx, migration, migration.Down); err != nil {
			return fmt.Errorf("rollback %s_%s failed: %w", migration.Version, migration.Name, err)
		}
		return r.forget(ctx, migration.Version)
	}
	log.Println("[migrate] nothing to roll back")
	return nil
}

// Status reports which migrations are applied and which are still pending.
func (r *Runner) Status(ctx context.Context) error {
	if err := r.ensureSchemaMigrations(ctx); err != nil {
		return err
	}
	applied, err := r.appliedVersions(ctx)
	if err != nil {
		return err
	}
	for _, migration := range Migrations() {
		state := "pending"
		if applied[migration.Version] {
			state = "up"
		}
		log.Printf("[migrate] %-7s %s_%s", state, migration.Version, migration.Name)
	}
	return nil
}

// record stores a version as applied.
func (r *Runner) record(ctx context.Context, version string) error {
	return r.db.WithContext(ctx).Exec(
		fmt.Sprintf("INSERT INTO %s (version) VALUES (?)", schemaMigrationsTable), version,
	).Error
}

// forget removes a version from the ledger after a rollback.
func (r *Runner) forget(ctx context.Context, version string) error {
	return r.db.WithContext(ctx).Exec(
		fmt.Sprintf("DELETE FROM %s WHERE version = ?", schemaMigrationsTable), version,
	).Error
}

// run executes one migration inside a transaction so a failure leaves the
// database untouched.
func (r *Runner) run(ctx context.Context, migration Migration, fn func(tx *gorm.DB) error) error {
	if fn == nil {
		return nil
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
