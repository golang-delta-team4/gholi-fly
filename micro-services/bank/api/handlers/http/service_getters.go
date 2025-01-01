package http

import (
	"context"

	"gholi-fly-bank/api/service"
	"gholi-fly-bank/app"
)

// Wallet service transient instance handler
func walletServiceGetter(appContainer app.App) ServiceGetter[*service.WalletService] {
	return func(ctx context.Context) *service.WalletService {
		return service.NewWalletService(appContainer.WalletService(ctx), appContainer.TransactionService(ctx))
	}
}
