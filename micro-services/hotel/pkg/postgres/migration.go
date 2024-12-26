package postgres

import (
	"fmt"
	"gholi-fly-hotel/pkg/adapters/storage/types"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := addUUIDExtension(db); err != nil {
		return fmt.Errorf("failed to add UUID extension: %w", err)
	}

	if err := db.AutoMigrate(
		&types.Hotel{},
		&types.Staff{},
		&types.Room{},
		&types.Booking{},
		// &types.BookingRoom{},
		// &types.Invoice{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}

func addUUIDExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
}
