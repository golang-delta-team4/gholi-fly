package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Wallet represents the GORM entity for the Wallet domain.
type Wallet struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID      uuid.UUID      `gorm:"not null"`                                        // Reference to the owner of the wallet.
	Type         uint8          `gorm:"not null;default:0"`                              // Wallet type (convertible to domain.WalletType).
	Balance      uint           `gorm:"not null;default:0"`                              // Current balance of the wallet.
	CreatedAt    time.Time      `gorm:"autoCreateTime"`                                  // Automatically set at creation.
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`                                  // Automatically updated.
	DeletedAt    gorm.DeletedAt `gorm:"index"`                                           // Soft delete support.
	Transactions []Transaction  `gorm:"foreignKey:WalletID;constraint:OnDelete:CASCADE"` // One-to-many relation with transactions.
	CreditCards  []*CreditCard  `gorm:"many2many:wallet_credit_cards"`                   // Many-to-many relation with credit cards.
}
