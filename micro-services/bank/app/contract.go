package app

import (
	"context"

	"gholi-fly-bank/config"
	creditPort "gholi-fly-bank/internal/credit/port"
	factorPort "gholi-fly-bank/internal/factor/port"
	transactionPort "gholi-fly-bank/internal/transaction/port"
	walletPort "gholi-fly-bank/internal/wallet/port"

	"gorm.io/gorm"
)

type App interface {
	WalletService(ctx context.Context) walletPort.Service
	TransactionService(ctx context.Context) transactionPort.Service
	CreditService(ctx context.Context) creditPort.Service
	FactorService(ctx context.Context) factorPort.Service
	DB() *gorm.DB
	Config() config.Config
}
