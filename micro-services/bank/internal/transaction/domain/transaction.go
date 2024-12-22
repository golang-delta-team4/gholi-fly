package domain

import (
	"time"

	"github.com/google/uuid"
)

// TransactionUUID is a UUID used for external references.
type TransactionUUID = uuid.UUID

// TransactionType defines the type of transaction.
type TransactionType uint8

const (
	TransactionTypeUnknown TransactionType = iota
	TransactionTypeCredit                  // Funds added to a wallet.
	TransactionTypeDebit                   // Funds removed from a wallet.
)

// TransactionStatus defines the status of a transaction.
type TransactionStatus uint8

const (
	TransactionStatusUnknown TransactionStatus = iota
	TransactionStatusPending
	TransactionStatusCompleted
	TransactionStatusFailed
	TransactionStatusRefunded
)

// Transaction represents a money movement in the system.
type Transaction struct {
	ID          TransactionUUID
	WalletID    uint              // ID of the wallet involved in the transaction.
	FactorID    uint              // Associated factor ID (if applicable).
	Amount      uint              // Transaction amount.
	Type        TransactionType   // Credit or Debit.
	Status      TransactionStatus // Pending, Completed, etc.
	Description string            // Optional description of the transaction.
	CreatedAt   time.Time         // Timestamp when the transaction was created.
	UpdatedAt   time.Time         // Timestamp when the transaction was last updated.
}

// TransactionFilters are used to filter transaction queries.
type TransactionFilters struct {
	WalletID uuid.UUID         // Filter by wallet ID.
	FactorID uuid.UUID         // Filter by associated factor.
	Type     TransactionType   // Filter by transaction type.
	Status   TransactionStatus // Filter by transaction status.
	DateFrom *time.Time        // Filter by starting date.
	DateTo   *time.Time        // Filter by ending date.
}
