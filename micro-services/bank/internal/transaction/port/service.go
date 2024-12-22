package port

import (
	"context"
	"gholi-fly-bank/internal/transaction/domain"
)

type Service interface {
	// Create a new transaction in the system.
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.TransactionUUID, error)

	// Retrieve a transaction by its ID.
	GetTransactionByID(ctx context.Context, transactionID domain.TransactionUUID) (*domain.Transaction, error)

	// Retrieve transactions based on filters.
	GetTransactions(ctx context.Context, filters domain.TransactionFilters) ([]domain.Transaction, error)

	// Update the status of a transaction.
	UpdateTransactionStatus(ctx context.Context, transactionID domain.TransactionUUID, status domain.TransactionStatus) error
}
