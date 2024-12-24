package postgres

import (
	"fmt"
	"gholi-fly-bank/pkg/adapters/storage/types"
	"log"

	"github.com/google/uuid"
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
		&types.Wallet{},
		&types.CreditCard{},
		&types.WalletCreditCard{},
		&types.Transaction{},
		&types.Factor{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// Ensure central wallet exists
	if err := ensureCentralWallet(db); err != nil {
		return fmt.Errorf("failed to ensure central wallet: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}

// addUUIDExtension ensures the UUID extension is added to the database.
func addUUIDExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public`).Error
}

// ensureCentralWallet creates a central wallet if it doesn't already exist.
func ensureCentralWallet(db *gorm.DB) error {
	centralWalletID := uuid.MustParse("11111111-1111-1111-1111-111111111111") // UUID of all ones

	var count int64
	err := db.Table("wallets").Where("id = ?", centralWalletID).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to check central wallet existence: %w", err)
	}

	if count == 0 {
		centralWallet := types.Wallet{
			ID:      centralWalletID,
			OwnerID: centralWalletID,
			Type:    3, // Assuming central wallet is of type company
			Balance: 0,
		}

		if err := db.Create(&centralWallet).Error; err != nil {
			return fmt.Errorf("failed to create central wallet: %w", err)
		}
		log.Println("Central wallet created successfully.")
	} else {
		log.Println("Central wallet already exists.")
	}

	return nil
}
