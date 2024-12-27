package postgres

import (
	"fmt"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"log"

	"gorm.io/gorm"
)

// Migrate applies all database migrations.
func Migrate(db *gorm.DB) error {
	// Add the UUID extension if not exists
	if err := addUUIDExtension(db); err != nil {
		return fmt.Errorf("failed to add UUID extension: %w", err)
	}

	// AutoMigrate models
	if err := db.AutoMigrate(
		// List all GORM models here
		&types.Agency{},
		&types.Staff{},
		&types.Factor{},
		&types.Tour{},
		&types.Reservation{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}

// addUUIDExtension ensures the UUID extension is added to the database.
func addUUIDExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public`).Error
}
