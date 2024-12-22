package storage

import (
	"context"
	"errors"
	"gholi-fly-bank/internal/wallet/domain"
	"gholi-fly-bank/internal/wallet/port"
	"gholi-fly-bank/pkg/adapters/storage/mapper"
	"gholi-fly-bank/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type walletRepo struct {
	db *gorm.DB
}

// NewWalletRepo creates a new instance of the wallet repository.
func NewWalletRepo(db *gorm.DB) port.Repo {
	return &walletRepo{db: db}
}

func (r *walletRepo) Create(ctx context.Context, walletDomain domain.Wallet) (domain.WalletUUID, error) {
	wallet := mapper.WalletDomain2Storage(walletDomain)
	err := r.db.WithContext(ctx).Table("wallets").Create(wallet).Error
	if err != nil {
		return domain.WalletUUID{}, err
	}
	return domain.WalletUUID(wallet.ID), nil
}

func (r *walletRepo) GetByID(ctx context.Context, walletID domain.WalletUUID) (*domain.Wallet, error) {
	var wallet types.Wallet
	err := r.db.WithContext(ctx).Table("wallets").Where("id = ?", walletID).First(&wallet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.WalletStorage2Domain(wallet), nil
}

func (r *walletRepo) Get(ctx context.Context, filters domain.WalletFilters) ([]domain.Wallet, error) {
	var wallets []types.Wallet
	query := r.db.WithContext(ctx).Table("wallets")

	// Apply filters
	if filters.OwnerID != uuid.Nil {
		query = query.Where("owner_id = ?", filters.OwnerID)
	}
	if filters.Type > 0 {
		query = query.Where("type = ?", uint8(filters.Type))
	}

	err := query.Find(&wallets).Error
	if err != nil {
		return nil, err
	}

	return mapper.BatchWalletStorage2Domain(wallets), nil
}

func (r *walletRepo) UpdateBalance(ctx context.Context, walletID domain.WalletUUID, newBalance uint) error {
	result := r.db.WithContext(ctx).Table("wallets").
		Where("id = ?", walletID).
		Update("balance", newBalance)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("wallet not found")
	}

	return nil
}
