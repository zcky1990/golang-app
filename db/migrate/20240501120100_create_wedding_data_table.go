package migrate

import "gorm.io/gorm"

// Migrated from the Mongo "wedding_data" collection. Nested payloads that were
// embedded documents are now jsonb columns, which keeps the API shape identical.
func init() {
	Register(Migration{
		Version: "20240501120100",
		Name:    "create_wedding_data_table",
		Up: func(tx *gorm.DB) error {
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS wedding_data (
					id uuid PRIMARY KEY,
					user_id uuid,
					bride jsonb,
					groom jsonb,
					wedding jsonb,
					love_story jsonb,
					envelop jsonb,
					gallery jsonb,
					guest jsonb,
					streaming jsonb,
					gifts jsonb,
					created_at timestamptz NOT NULL DEFAULT now(),
					updated_at timestamptz NOT NULL DEFAULT now()
				)`).Error; err != nil {
				return err
			}
			return tx.Exec(`
				CREATE INDEX IF NOT EXISTS index_wedding_data_on_user_id ON wedding_data (user_id)`).Error
		},
		Down: func(tx *gorm.DB) error {
			return tx.Exec(`DROP TABLE IF EXISTS wedding_data`).Error
		},
	})
}
