package port

import (
	"context"
	"gholi-fly-bank/internal/transaction/domain"
)

type Repo interface {
	// Create a new transaction in the system.
	Create(ctx context.Context, transaction domain.Transaction) (domain.TransactionUUID, error)

	// Retrieve a transaction by its ID.
	GetByID(ctx context.Context, transactionID domain.TransactionUUID) (*domain.Transaction, error)

	// Retrieve transactions based on filters.
	Get(ctx context.Context, filters domain.TransactionFilters) ([]domain.Transaction, error)

	// Update the status of a transaction.
	UpdateStatus(ctx context.Context, transactionID domain.TransactionUUID, status domain.TransactionStatus) error

	GetSum(ctx context.Context, filters domain.TransactionFilters) (int64, error)
}
