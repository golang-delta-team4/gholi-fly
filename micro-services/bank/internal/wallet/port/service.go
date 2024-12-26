package port

import (
	"context"
	"gholi-fly-bank/internal/wallet/domain"
)

type Service interface {
	// Create a wallet in the system.
	CreateWallet(ctx context.Context, wallet domain.Wallet) (domain.WalletUUID, error)

	// Retrieve a wallet by its ID.
	GetWalletByID(ctx context.Context, walletID domain.WalletUUID) (*domain.Wallet, error)

	// Retrieve wallets based on filters.
	GetWallets(ctx context.Context, filters domain.WalletFilters) ([]domain.Wallet, error)

	// Update the balance of a wallet.
	UpdateWalletBalance(ctx context.Context, walletID domain.WalletUUID, newBalance uint) error
}
