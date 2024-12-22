package handlers

import (
	"context"
	"gholi-fly-bank/api/pb"
	"gholi-fly-bank/app"
	creditPort "gholi-fly-bank/internal/credit/port"
	factorPort "gholi-fly-bank/internal/factor/port"
	transactionPort "gholi-fly-bank/internal/transaction/port"
	"gholi-fly-bank/internal/wallet/domain"
	walletPort "gholi-fly-bank/internal/wallet/port"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCBankHandler implements all gRPC endpoints for the bank service.
type GRPCBankHandler struct {
	pb.UnimplementedWalletServiceServer
	walletService      walletPort.Service
	transactionService transactionPort.Service
	creditService      creditPort.Service
	factorService      factorPort.Service
}

// NewGRPCBankHandler initializes a new GRPCBankHandler with all necessary services.
func NewGRPCBankHandler(ctx context.Context, app app.App) *GRPCBankHandler {
	return &GRPCBankHandler{
		walletService:      app.WalletService(ctx),
		transactionService: app.TransactionService(ctx),
		creditService:      app.CreditService(ctx),
		factorService:      app.FactorService(ctx),
	}
}

// CreateWallet handles the CreateWallet gRPC method.
func (h *GRPCBankHandler) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	// Parse OwnerId from string to uuid.UUID
	ownerID, err := uuid.Parse(req.OwnerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid owner ID format: %v", err)
	}

	// Check if a wallet with the given owner already exists
	existingWallets, err := h.walletService.GetWallets(ctx, domain.WalletFilters{OwnerID: ownerID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check existing wallets: %v", err)
	}

	if len(existingWallets) > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "A wallet for this owner already exists.")
	}

	// Create a new wallet
	wallet := domain.Wallet{
		OwnerID: ownerID,
		Type:    domain.WalletType(req.Type),
		Balance: 0, // Default balance
	}

	walletID, err := h.walletService.CreateWallet(ctx, wallet)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create wallet: %v", err)
	}

	return &pb.CreateWalletResponse{
		Status: pb.ResponseStatus_SUCCESS,
		Wallet: &pb.Wallet{
			Id:      walletID.String(),
			OwnerId: wallet.OwnerID.String(),
			Type:    pb.WalletType(wallet.Type),
			Balance: uint64(wallet.Balance),
		},
	}, nil
}

// GetWallets handles the GetWallets gRPC method.
func (h *GRPCBankHandler) GetWallets(ctx context.Context, req *pb.GetWalletsRequest) (*pb.GetWalletsResponse, error) {
	// Parse OwnerId from string to uuid.UUID
	var ownerID uuid.UUID
	if req.OwnerId != "" {
		parsedOwnerID, err := uuid.Parse(req.OwnerId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid owner ID format: %v", err)
		}
		ownerID = parsedOwnerID
	}

	// Convert request filters to domain filters
	filters := domain.WalletFilters{
		OwnerID: ownerID,
		Type:    domain.WalletType(req.Type),
	}

	// Call wallet service to fetch wallets
	wallets, err := h.walletService.GetWallets(ctx, filters)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve wallets: %v", err)
	}

	// Convert domain wallets to protobuf Wallet messages
	var walletResponses []*pb.Wallet
	for _, w := range wallets {
		walletResponses = append(walletResponses, &pb.Wallet{
			Id:        w.ID.String(),
			OwnerId:   w.OwnerID.String(), // Convert uuid.UUID to string
			Type:      pb.WalletType(w.Type),
			Balance:   uint64(w.Balance),
			CreatedAt: w.CreatedAt.Format("2006-01-02T15:04:05Z07:00"), // Convert time to RFC3339
			UpdatedAt: w.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &pb.GetWalletsResponse{
		Status:  pb.ResponseStatus_SUCCESS,
		Wallets: walletResponses,
	}, nil
}
