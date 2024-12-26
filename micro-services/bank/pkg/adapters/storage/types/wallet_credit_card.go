package types

import "time"

// WalletCreditCard represents the association table for Wallet and CreditCard.
type WalletCreditCard struct {
	WalletID     uint      `gorm:"primaryKey"`
	CreditCardID uint      `gorm:"primaryKey"`
	LinkedAt     time.Time `gorm:"autoCreateTime"` // Timestamp when the link was established.
}
