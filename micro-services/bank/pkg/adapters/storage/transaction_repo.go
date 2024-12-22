package storage

import (
	"context"
	"errors"
	"gholi-fly-bank/internal/transaction/domain"
	"gholi-fly-bank/internal/transaction/port"
	"gholi-fly-bank/pkg/adapters/storage/mapper"
	"gholi-fly-bank/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

// NewTransactionRepo creates a new instance of the transaction repository.
func NewTransactionRepo(db *gorm.DB) port.Repo {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) Create(ctx context.Context, transactionDomain domain.Transaction) (domain.TransactionUUID, error) {
	transaction := mapper.TransactionDomain2Storage(transactionDomain)
	err := r.db.WithContext(ctx).Table("transactions").Create(transaction).Error
	if err != nil {
		return domain.TransactionUUID{}, err
	}
	return domain.TransactionUUID(transaction.ID), nil
}

func (r *transactionRepo) GetByID(ctx context.Context, transactionID domain.TransactionUUID) (*domain.Transaction, error) {
	var transaction types.Transaction
	err := r.db.WithContext(ctx).Table("transactions").Where("id = ?", transactionID).First(&transaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.TransactionStorage2Domain(transaction), nil
}

func (r *transactionRepo) Get(ctx context.Context, filters domain.TransactionFilters) ([]domain.Transaction, error) {
	var transactions []types.Transaction
	query := r.db.WithContext(ctx).Table("transactions")

	// Apply filters
	if filters.WalletID != uuid.Nil {
		query = query.Where("wallet_id = ?", filters.WalletID)
	}
	if filters.FactorID != uuid.Nil {
		query = query.Where("factor_id = ?", filters.FactorID)
	}
	if filters.Type > 0 {
		query = query.Where("type = ?", uint8(filters.Type))
	}
	if filters.Status > 0 {
		query = query.Where("status = ?", uint8(filters.Status))
	}
	if filters.DateFrom != nil {
		query = query.Where("created_at >= ?", *filters.DateFrom)
	}
	if filters.DateTo != nil {
		query = query.Where("created_at <= ?", *filters.DateTo)
	}

	err := query.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return mapper.BatchTransactionStorage2Domain(transactions), nil
}

func (r *transactionRepo) UpdateStatus(ctx context.Context, transactionID domain.TransactionUUID, status domain.TransactionStatus) error {
	result := r.db.WithContext(ctx).Table("transactions").
		Where("id = ?", transactionID).
		Update("status", uint8(status))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}

	return nil
}
