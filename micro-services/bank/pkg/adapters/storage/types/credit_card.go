package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreditCard represents the GORM entity for the CreditCard domain.
type CreditCard struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CardNumber string         `gorm:"type:char(16);not null;uniqueIndex"` // 16-digit Iranian bank card number.
	ExpiryDate string         `gorm:"type:char(5);not null"`              // Expiration date in MM/YY format.
	CVV        string         `gorm:"type:char(3);not null"`              // Security code.
	HolderName string         `gorm:"type:varchar(255);not null"`         // Cardholder's name.
	CreatedAt  time.Time      `gorm:"autoCreateTime"`                     // Automatically set at creation.
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`                     // Automatically updated.
	DeletedAt  gorm.DeletedAt `gorm:"index"`                              // Soft delete support.
	Wallets    []*Wallet      `gorm:"many2many:wallet_credit_cards"`      // Many-to-many relation with wallets.
}
