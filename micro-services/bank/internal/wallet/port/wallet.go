package port

import (
	"context"
	"gholi-fly-bank/internal/wallet/domain"
)

type Repo interface {
	// Create a new wallet in the system.
	Create(ctx context.Context, wallet domain.Wallet) (domain.WalletUUID, error)

	// Retrieve a wallet by its ID.
	GetByID(ctx context.Context, walletID domain.WalletUUID) (*domain.Wallet, error)

	// Retrieve wallets based on filters.
	Get(ctx context.Context, filters domain.WalletFilters) ([]domain.Wallet, error)

	// Update the balance of a wallet.
	UpdateBalance(ctx context.Context, walletID domain.WalletUUID, newBalance uint) error
}
