package transaction

import (
	"context"
	"errors"
	"fmt"
	"gholi-fly-bank/internal/transaction/domain"
	"gholi-fly-bank/internal/transaction/port"
	"log"
)

var (
	ErrTransactionCreation      = errors.New("error creating transaction")
	ErrTransactionValidation    = errors.New("transaction validation failed")
	ErrTransactionNotFound      = errors.New("transaction not found")
	ErrTransactionStatusUpdate  = errors.New("error updating transaction status")
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
)

type service struct {
	repo port.Repo
}

// NewService creates a new instance of the transaction service.
func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.TransactionUUID, error) {
	// Validate transaction details
	if transaction.Amount <= 0 {
		return domain.TransactionUUID{}, fmt.Errorf("%w: amount must be greater than zero", ErrInvalidTransactionAmount)
	}
	if transaction.Type == domain.TransactionTypeUnknown {
		return domain.TransactionUUID{}, fmt.Errorf("%w: invalid transaction type", ErrTransactionValidation)
	}

	transactionID, err := s.repo.Create(ctx, transaction)
	if err != nil {
		log.Println("error creating transaction:", err.Error())
		return domain.TransactionUUID{}, ErrTransactionCreation
	}

	return transactionID, nil
}

func (s *service) GetTransactionByID(ctx context.Context, transactionID domain.TransactionUUID) (*domain.Transaction, error) {
	transaction, err := s.repo.GetByID(ctx, transactionID)
	if err != nil {
		log.Println("error fetching transaction by ID:", err.Error())
		return nil, err
	}

	if transaction == nil {
		return nil, ErrTransactionNotFound
	}

	return transaction, nil
}

func (s *service) GetTransactions(ctx context.Context, filters domain.TransactionFilters) ([]domain.Transaction, error) {
	transactions, err := s.repo.Get(ctx, filters)
	if err != nil {
		log.Println("error fetching transactions:", err.Error())
		return nil, err
	}

	return transactions, nil
}

func (s *service) UpdateTransactionStatus(ctx context.Context, transactionID domain.TransactionUUID, status domain.TransactionStatus) error {
	// Validate the status transition if necessary
	if status == domain.TransactionStatusUnknown {
		return fmt.Errorf("%w: invalid transaction status update", ErrTransactionStatusUpdate)
	}

	err := s.repo.UpdateStatus(ctx, transactionID, status)
	if err != nil {
		log.Println("error updating transaction status:", err.Error())
		return ErrTransactionStatusUpdate
	}

	return nil
}
