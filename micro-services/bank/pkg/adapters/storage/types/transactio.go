package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Transaction represents the GORM entity for the Transaction domain.
type Transaction struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	WalletID    uuid.UUID      `gorm:"not null;index"`                                   // Foreign key referencing Wallet.
	Wallet      Wallet         `gorm:"foreignKey:WalletID;constraint:OnDelete:CASCADE"`  // Relation to Wallet.
	FactorID    *uuid.UUID     `gorm:"index"`                                            // Foreign key referencing Factor (nullable).
	Factor      *Factor        `gorm:"foreignKey:FactorID;constraint:OnDelete:SET NULL"` // Relation to Factor.
	Amount      uint           `gorm:"not null"`                                         // Transaction amount.
	Type        uint8          `gorm:"not null;default:0"`                               // Transaction type (0: Unknown, 1: Credit, 2: Debit).
	Status      uint8          `gorm:"not null;default:0"`                               // Transaction status (0: Unknown, 1: Pending, 2: Completed, etc.).
	Description string         `gorm:"type:varchar(255)"`                                // Optional description of the transaction.
	CreatedAt   time.Time      `gorm:"autoCreateTime"`                                   // Automatically set at creation.
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`                                   // Automatically updated.
	DeletedAt   gorm.DeletedAt `gorm:"index"`                                            // Soft delete support.
}
