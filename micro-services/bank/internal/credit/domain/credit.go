package domain

import (
	"time"

	"github.com/google/uuid"
)

// CreditCardUUID is a UUID used for external references.
type CreditCardUUID = uuid.UUID

// CreditCard represents a linked credit card in the system.
type CreditCard struct {
	ID         CreditCardUUID
	CardNumber string    // 16-digit Iranian bank card number.
	ExpiryDate string    // Expiration date in MM/YY format.
	CVV        string    // Security code.
	HolderName string    // Name of the cardholder.
	CreatedAt  time.Time // Timestamp when the card was added.
	UpdatedAt  time.Time // Timestamp when the card details were last updated.
}

// CreditCardFilters are used to filter credit card queries.
type CreditCardFilters struct {
	CardNumber string // Search by card number.
	HolderName string // Search by holder's name.
}
