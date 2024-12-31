package service

import (
	"context"
	"errors"
	pb "gholi-fly-bank/api/pb"
	transactionDomain "gholi-fly-bank/internal/transaction/domain"
	transactionPort "gholi-fly-bank/internal/transaction/port"
	"gholi-fly-bank/internal/wallet/domain"
	"gholi-fly-bank/internal/wallet/port"
	"log"

	"github.com/google/uuid"
)

var (
	ErrWalletCreation           = errors.New("error on creating new wallet")
	ErrWalletValidation         = errors.New("wallet validation failed")
	ErrWalletNotFound           = errors.New("wallet not found")
	ErrBalanceUpdate            = errors.New("error updating wallet balance")
	ErrInsufficientBalance      = errors.New("insufficient balance")
	ErrUnauthorizedWalletAccess = errors.New("unauthorized access to wallet")
)

type WalletService struct {
	walletRepo      port.Service
	transactionRepo transactionPort.Service
}

// NewWalletService creates a new instance of WalletService.
func NewWalletService(walletRepo port.Service, trasactionRepo transactionPort.Service) *WalletService {
	return &WalletService{
		walletRepo:      walletRepo,
		transactionRepo: trasactionRepo,
	}
}

func (ws *WalletService) GetWallets(ctx context.Context, ownerID uuid.UUID) (*pb.GetWalletsResponse, error) {

	filters := domain.WalletFilters{
		OwnerID: ownerID,
	}

	// Retrieve wallets based on filters
	wallets, err := ws.walletRepo.GetWallets(ctx, filters)
	if err != nil {
		log.Println("error fetching wallets:", err)
		return nil, err
	}

	var pbWallets []*pb.Wallet
	for _, w := range wallets {
		// Calculate the pending credit transactions
		creditFilters := transactionDomain.TransactionFilters{
			WalletID: w.ID,
			Type:     transactionDomain.TransactionTypeCredit,
			Status:   transactionDomain.TransactionStatusPending,
		}
		creditSum, err := ws.transactionRepo.GetTransactionSum(ctx, &creditFilters)
		if err != nil {
			log.Println("error fetching pending credit transactions:", err)
			return nil, err
		}

		// Calculate the pending debit transactions
		debitFilters := transactionDomain.TransactionFilters{
			WalletID: w.ID,
			Type:     transactionDomain.TransactionTypeDebit,
			Status:   transactionDomain.TransactionStatusPending,
		}
		debitSum, err := ws.transactionRepo.GetTransactionSum(ctx, &debitFilters)
		if err != nil {
			log.Println("error fetching pending debit transactions:", err)
			return nil, err
		}

		// Calculate the actual balance
		actualBalance := int64(w.Balance) + creditSum - debitSum

		// Append to response
		pbWallets = append(pbWallets, &pb.Wallet{
			Id:        w.ID.String(),
			OwnerId:   w.OwnerID.String(),
			Type:      pb.WalletType(w.Type),
			Balance:   uint64(actualBalance), // Updated balance calculation
			CreatedAt: w.CreatedAt.String(),
			UpdatedAt: w.UpdatedAt.String(),
		})
	}

	return &pb.GetWalletsResponse{
		Status:  pb.ResponseStatus_SUCCESS,
		Wallets: pbWallets,
	}, nil
}

func (ws *WalletService) UpdateWalletBalance(ctx context.Context, walletID string, userUUID uuid.UUID, depositAmount uint64) error {
	id, err := uuid.Parse(walletID)
	if err != nil {
		log.Println("invalid wallet ID:", err)
		return ErrWalletValidation
	}

	wallet, err := ws.walletRepo.GetWalletByID(ctx, domain.WalletUUID(id))
	if err != nil {
		log.Println("error fetching wallet for balance update:", err)
		return ErrBalanceUpdate
	}

	if wallet == nil {
		return ErrWalletNotFound
	}
	if wallet.OwnerID != userUUID {
		return ErrUnauthorizedWalletAccess
	}
	newBalance := wallet.Balance + uint(depositAmount)
	err = ws.walletRepo.UpdateWalletBalance(ctx, domain.WalletUUID(id), uint(newBalance))
	if err != nil {
		log.Println("error updating wallet balance:", err)
		return ErrBalanceUpdate
	}

	return nil
}
