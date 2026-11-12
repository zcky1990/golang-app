// Package postgres is the PostgreSQL counterpart of config/mongo.
//
// It opens the primary database connection using GORM (the closest thing Go
// has to ActiveRecord) and exposes the *gorm.DB handle to the rest of the
// application. The old Mongo package is intentionally kept untouched so the
// project can still be pointed at MongoDB while the switch-over is finished.
package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	c "golang_app/golangApp/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresDB mirrors the MongoDB struct in config/mongo.
type PostgresDB struct {
	DB *gorm.DB
}

// NewPostgres connects to PostgreSQL using the POSTGRES_* environment
// variables and returns the wrapped GORM handle.
func NewPostgres() (*PostgresDB, error) {
	dsn := BuildDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode()),
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting PostgreSQL connection pool: %w", err)
	}
	sqlDB.SetMaxOpenConns(envInt(c.POSTGRES_MAX_CONNS, 25))
	sqlDB.SetMaxIdleConns(envInt(c.POSTGRES_MAX_IDLE_CONNS, 5))
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging PostgreSQL: %w", err)
	}

	log.Println("Connected to PostgreSQL!")
	return &PostgresDB{DB: db}, nil
}

// BuildDSN builds a libpq keyword/value connection string from the
// environment, the same way config/mongo builds its Mongo URL.
func BuildDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		envOr(c.POSTGRES_HOST, "localhost"),
		envOr(c.POSTGRES_PORT, "5432"),
		envOr(c.POSTGRES_USER, "postgres"),
		os.Getenv(c.POSTGRES_PASSWORD),
		envOr(c.POSTGRES_DB, "golang_app_development"),
		envOr(c.POSTGRES_SSLMODE, "disable"),
		envOr(c.POSTGRES_TIMEZONE, "UTC"),
	)
}

// Close releases the underlying connection pool.
func (p *PostgresDB) Close() error {
	if p == nil || p.DB == nil {
		return nil
	}
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error disconnecting from PostgreSQL: %v", err)
		return err
	}
	log.Println("Disconnected from PostgreSQL")
	return nil
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

// logMode keeps query logging quiet in production and verbose everywhere else,
// similar to Rails' log_level per environment.
func logMode() logger.LogLevel {
	if os.Getenv("ENV") == "production" {
		return logger.Warn
	}
	return logger.Info
}
