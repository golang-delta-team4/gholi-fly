package wallet

import (
	"context"
	"errors"
	"gholi-fly-bank/internal/wallet/domain"
	"gholi-fly-bank/internal/wallet/port"
	"log"
)

var (
	ErrWalletCreation      = errors.New("error on creating new wallet")
	ErrWalletValidation    = errors.New("wallet validation failed")
	ErrWalletNotFound      = errors.New("wallet not found")
	ErrBalanceUpdate       = errors.New("error updating wallet balance")
	ErrInsufficientBalance = errors.New("insufficient balance")
)

type service struct {
	repo port.Repo
}

// NewService creates a new instance of the wallet service.
func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateWallet(ctx context.Context, wallet domain.Wallet) (domain.WalletUUID, error) {

	walletID, err := s.repo.Create(ctx, wallet)
	if err != nil {
		log.Println("error on creating new wallet:", err.Error())
		return domain.WalletUUID{}, ErrWalletCreation
	}

	return walletID, nil
}

func (s *service) GetWalletByID(ctx context.Context, walletID domain.WalletUUID) (*domain.Wallet, error) {
	wallet, err := s.repo.GetByID(ctx, walletID)
	if err != nil {
		log.Println("error fetching wallet by ID:", err.Error())
		return nil, err
	}

	if wallet == nil {
		return nil, ErrWalletNotFound
	}

	return wallet, nil
}

func (s *service) GetWallets(ctx context.Context, filters domain.WalletFilters) ([]domain.Wallet, error) {
	wallets, err := s.repo.Get(ctx, filters)
	if err != nil {
		log.Println("error fetching wallets with filters:", err.Error())
		return nil, err
	}

	return wallets, nil
}

func (s *service) UpdateWalletBalance(ctx context.Context, walletID domain.WalletUUID, newBalance uint) error {
	// Fetch the wallet to check existing balance
	wallet, err := s.repo.GetByID(ctx, walletID)
	if err != nil {
		log.Println("error fetching wallet for balance update:", err.Error())
		return ErrBalanceUpdate
	}

	if wallet == nil {
		return ErrWalletNotFound
	}

	err = s.repo.UpdateBalance(ctx, walletID, newBalance)
	if err != nil {
		log.Println("error updating wallet balance:", err.Error())
		return ErrBalanceUpdate
	}

	return nil
}
