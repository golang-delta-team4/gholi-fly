package mapper

import (
	"gholi-fly-bank/internal/wallet/domain"
	"gholi-fly-bank/pkg/adapters/storage/types"
	"gholi-fly-bank/pkg/fp"
)

// WalletDomain2Storage converts a Wallet from the domain layer to the storage layer.
func WalletDomain2Storage(walletDomain domain.Wallet) *types.Wallet {
	return &types.Wallet{
		ID:        walletDomain.ID,
		OwnerID:   walletDomain.OwnerID,
		Type:      uint8(walletDomain.Type), // Convert domain.WalletType to uint8.
		Balance:   walletDomain.Balance,
		CreatedAt: walletDomain.CreatedAt,
		UpdatedAt: walletDomain.UpdatedAt,
	}
}
func walletDomain2Storage(walletDomain domain.Wallet) types.Wallet {
	return types.Wallet{
		ID:        walletDomain.ID,
		OwnerID:   walletDomain.OwnerID,
		Type:      uint8(walletDomain.Type), // Convert domain.WalletType to uint8.
		Balance:   walletDomain.Balance,
		CreatedAt: walletDomain.CreatedAt,
		UpdatedAt: walletDomain.UpdatedAt,
	}
}
func BatchWalletDomain2Storage(domains []domain.Wallet) []types.Wallet {
	return fp.Map(domains, walletDomain2Storage)
}

// WalletStorage2Domain converts a Wallet from the storage layer to the domain layer.
func WalletStorage2Domain(wallet types.Wallet) *domain.Wallet {
	return &domain.Wallet{
		ID:        wallet.ID,
		OwnerID:   wallet.OwnerID,
		Type:      domain.WalletType(wallet.Type), // Convert uint8 to domain.WalletType.
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}
}
func walletStorage2Domain(wallet types.Wallet) domain.Wallet {
	return domain.Wallet{
		ID:        wallet.ID,
		OwnerID:   wallet.OwnerID,
		Type:      domain.WalletType(wallet.Type), // Convert uint8 to domain.WalletType.
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}
}
func BatchWalletStorage2Domain(wallets []types.Wallet) []domain.Wallet {
	return fp.Map(wallets, walletStorage2Domain)
}
