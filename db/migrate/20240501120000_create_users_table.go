package migrate

import "gorm.io/gorm"

// Migrated from the Mongo "User" collection to the Rails-style users table.
func init() {
	Register(Migration{
		Version: "20240501120000",
		Name:    "create_users_table",
		Up: func(tx *gorm.DB) error {
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS users (
					id uuid PRIMARY KEY,
					username varchar(255),
					email varchar(255) NOT NULL,
					firstname varchar(255),
					lastname varchar(255),
					auth_token varchar(255),
					password varchar(255),
					created_at timestamptz NOT NULL DEFAULT now(),
					updated_at timestamptz NOT NULL DEFAULT now()
				)`).Error; err != nil {
				return err
			}
			// Same uniqueness guarantee the Mongo index used to provide.
			return tx.Exec(`
				CREATE UNIQUE INDEX IF NOT EXISTS index_users_on_email ON users (email)`).Error
		},
		Down: func(tx *gorm.DB) error {
			return tx.Exec(`DROP TABLE IF EXISTS users`).Error
		},
	})
}
