package transaction

import (
	"context"
	"errors"
	"fmt"
	"gholi-fly-bank/internal/transaction/domain"
	transactionRepo "gholi-fly-bank/internal/transaction/port"
	walletRepo "gholi-fly-bank/internal/wallet/port"
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
	transactionRepo transactionRepo.Repo
	walletRepo      walletRepo.Repo
}

// NewService creates a new instance of the transaction service.
func NewService(transactionRepo transactionRepo.Repo, walletRepo walletRepo.Repo) transactionRepo.Service {
	return &service{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
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

	// Check if the transaction status is `Success`
	if transaction.Status == domain.TransactionStatusCompleted {
		// Fetch the associated wallet
		wallet, err := s.walletRepo.GetByID(ctx, transaction.WalletID)
		if err != nil {
			log.Printf("error fetching wallet for transaction: %v", err)
			return domain.TransactionUUID{}, fmt.Errorf("failed to fetch wallet: %w", err)
		}
		if wallet == nil {
			return domain.TransactionUUID{}, fmt.Errorf("wallet not found for transaction: %s", transaction.WalletID)
		}

		// Ensure sufficient funds for debit transactions
		if transaction.Type == domain.TransactionTypeDebit && wallet.Balance < transaction.Amount {
			return domain.TransactionUUID{}, fmt.Errorf("%w: insufficient funds for wallet %s", ErrTransactionValidation, wallet.ID)
		}

		// Update wallet balance
		var newBalance uint
		switch transaction.Type {
		case domain.TransactionTypeCredit:
			newBalance = wallet.Balance + transaction.Amount
		case domain.TransactionTypeDebit:
			newBalance = wallet.Balance - transaction.Amount
		}

		err = s.walletRepo.UpdateBalance(ctx, wallet.ID, newBalance)
		if err != nil {
			log.Printf("error updating wallet balance: %v", err)
			return domain.TransactionUUID{}, fmt.Errorf("failed to update wallet balance: %w", err)
		}
	}

	// Create the transaction
	transactionID, err := s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		log.Printf("error creating transaction: %v", err)
		return domain.TransactionUUID{}, ErrTransactionCreation
	}

	return transactionID, nil
}

func (s *service) GetTransactionByID(ctx context.Context, transactionID domain.TransactionUUID) (*domain.Transaction, error) {
	transaction, err := s.transactionRepo.GetByID(ctx, transactionID)
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
	transactions, err := s.transactionRepo.Get(ctx, filters)
	if err != nil {
		log.Println("error fetching transactions:", err.Error())
		return nil, err
	}

	return transactions, nil
}

func (s *service) UpdateTransactionStatus(ctx context.Context, transactionID domain.TransactionUUID, status domain.TransactionStatus) error {
	// Validate the status transition
	if status == domain.TransactionStatusUnknown {
		return fmt.Errorf("%w: invalid transaction status update", ErrTransactionStatusUpdate)
	}

	// Fetch the transaction
	transaction, err := s.transactionRepo.GetByID(ctx, transactionID)
	if err != nil {
		log.Printf("error fetching transaction for ID %s: %v", transactionID, err)
		return ErrTransactionStatusUpdate
	}
	if transaction == nil {
		return fmt.Errorf("%w: transaction not found for ID %s", ErrTransactionNotFound, transactionID)
	}

	// Prevent completed transactions from being updated
	if transaction.Status == domain.TransactionStatusCompleted {
		return fmt.Errorf("%w: cannot change status of completed transaction", ErrTransactionStatusUpdate)
	}

	// Only update wallet balance if status is changing to success
	if status == domain.TransactionStatusCompleted && transaction.Status != domain.TransactionStatusCompleted {
		// Fetch the associated wallet
		wallet, err := s.walletRepo.GetByID(ctx, transaction.WalletID)
		if err != nil {
			log.Printf("error fetching wallet for transaction %s: %v", transactionID, err)
			return fmt.Errorf("failed to fetch wallet: %w", err)
		}
		if wallet == nil {
			return fmt.Errorf("associated wallet not found for transaction %s", transactionID)
		}

		var newBalance uint
		switch transaction.Type {
		case domain.TransactionTypeCredit:
			newBalance = wallet.Balance + transaction.Amount
		case domain.TransactionTypeDebit:
			// Ensure the wallet has sufficient funds for the debit
			if wallet.Balance < transaction.Amount {
				return fmt.Errorf("insufficient funds for wallet %s during debit transaction", wallet.ID)
			}
			newBalance = wallet.Balance - transaction.Amount
		default:
			log.Printf("invalid transaction type for ID %s", transactionID)
			return fmt.Errorf("%w: invalid transaction type", ErrTransactionValidation)
		}

		// Update wallet balance
		if err = s.walletRepo.UpdateBalance(ctx, wallet.ID, newBalance); err != nil {
			log.Printf("error updating wallet balance for wallet %s: %v", wallet.ID, err)
			return fmt.Errorf("failed to update wallet balance: %w", err)
		}
	}

	// Update the transaction's status
	if err = s.transactionRepo.UpdateStatus(ctx, transactionID, status); err != nil {
		log.Printf("error updating transaction status for ID %s: %v", transactionID, err)
		return ErrTransactionStatusUpdate
	}

	return nil
}

func (s *service) GetTransactionSum(ctx context.Context, filters *domain.TransactionFilters) (int64, error) {
	fmt.Printf("%+v", filters)
	total, err := s.transactionRepo.GetSum(ctx, *filters)
	if err != nil {
		log.Println("error fetching transaction sum:", err.Error())
		return 0, err
	}
	return total, nil
}
