package app

import (
	"context"
	"fmt"

	"gholi-fly-bank/config"
	"gholi-fly-bank/internal/credit"
	creditPort "gholi-fly-bank/internal/credit/port"
	"gholi-fly-bank/internal/factor"
	factorPort "gholi-fly-bank/internal/factor/port"
	"gholi-fly-bank/internal/transaction"
	transactionPort "gholi-fly-bank/internal/transaction/port"
	"gholi-fly-bank/internal/wallet"
	walletPort "gholi-fly-bank/internal/wallet/port"
	"gholi-fly-bank/pkg/adapters/storage"
	"gholi-fly-bank/pkg/postgres"

	"gorm.io/gorm"

	appCtx "gholi-fly-bank/pkg/context"
)

type app struct {
	db                 *gorm.DB
	cfg                config.Config
	walletService      walletPort.Service
	transactionService transactionPort.Service
	creditService      creditPort.Service
	factorService      factorPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

// WalletService provides the wallet service, with support for contextual DB switching.
func (a *app) WalletService(ctx context.Context) walletPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.walletService == nil {
			a.walletService = a.walletServiceWithDB(a.db)
		}
		return a.walletService
	}

	return a.walletServiceWithDB(db)
}

func (a *app) walletServiceWithDB(db *gorm.DB) walletPort.Service {
	return wallet.NewService(storage.NewWalletRepo(db))
}

// TransactionService provides the transaction service, with support for contextual DB switching.
func (a *app) TransactionService(ctx context.Context) transactionPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.transactionService == nil {
			a.transactionService = a.transactionServiceWithDB(a.db)
		}
		return a.transactionService
	}

	return a.transactionServiceWithDB(db)
}

func (a *app) transactionServiceWithDB(db *gorm.DB) transactionPort.Service {
	return transaction.NewService(storage.NewTransactionRepo(db), storage.NewWalletRepo(db))
}

// CreditService provides the credit card service, with support for contextual DB switching.
func (a *app) CreditService(ctx context.Context) creditPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.creditService == nil {
			a.creditService = a.creditServiceWithDB(a.db)
		}
		return a.creditService
	}

	return a.creditServiceWithDB(db)
}

func (a *app) creditServiceWithDB(db *gorm.DB) creditPort.Service {
	return credit.NewService(storage.NewCreditCardRepo(db))
}

// FactorService provides the factor service, with support for contextual DB switching.
func (a *app) FactorService(ctx context.Context) factorPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.factorService == nil {
			a.factorService = a.factorServiceWithDB(a.db)
		}
		return a.factorService
	}

	return a.factorServiceWithDB(db)
}

func (a *app) factorServiceWithDB(db *gorm.DB) factorPort.Service {
	return factor.NewService(storage.NewFactorRepo(db))
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Apply migrations
	if err := postgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	a.db = db
	return nil
}

// NewApp initializes a new App instance.
func NewApp(cfg config.Config) (App, error) {
	a := &app{cfg: cfg}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}

// NewMustApp initializes a new App instance and panics on error.
func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
