package domain

import (
	"time"

	"github.com/google/uuid"
)

// WalletType defines the type of wallet (e.g., Person, Company, App).
type WalletType uint8

const (
	WalletTypeUnknown WalletType = iota
	WalletTypePerson
	WalletTypeCompany
	WalletTypeApp
)

// WalletUUID is a UUID used for external references.
type WalletUUID = uuid.UUID

// Wallet represents a wallet in the system.
type Wallet struct {
	ID        WalletUUID
	OwnerID   uuid.UUID // Reference to the user, company, or app owning the wallet.
	Type      WalletType
	Balance   uint // The current balance of the wallet.
	CreatedAt time.Time
	UpdatedAt time.Time
}

// WalletFilters are used to filter wallet queries.
type WalletFilters struct {
	OwnerID uuid.UUID
	Type    WalletType
}
