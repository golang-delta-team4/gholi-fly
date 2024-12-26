package handlers

import (
	"context"
	"gholi-fly-bank/api/pb"
	"gholi-fly-bank/app"
	creditPort "gholi-fly-bank/internal/credit/port"
	factorDomain "gholi-fly-bank/internal/factor/domain"
	factorPort "gholi-fly-bank/internal/factor/port"

	transactionDomain "gholi-fly-bank/internal/transaction/domain"
	transactionPort "gholi-fly-bank/internal/transaction/port"
	"gholi-fly-bank/internal/wallet/domain"
	walletPort "gholi-fly-bank/internal/wallet/port"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCBankHandler implements all gRPC endpoints for the bank service.
type GRPCBankHandler struct {
	pb.UnimplementedWalletServiceServer
	pb.UnimplementedFactorServiceServer
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
func (h *GRPCBankHandler) CreateFactor(ctx context.Context, req *pb.CreateFactorRequest) (*pb.CreateFactorResponse, error) {
	// Validate required fields
	if req.Factor == nil {
		return nil, status.Errorf(codes.InvalidArgument, "factor is required")
	}

	if req.Factor.CustomerId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "customer_id is required")
	}

	customerID, err := uuid.Parse(req.Factor.CustomerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid customer_id format: %v", err)
	}

	// Map request to domain
	factor := factorDomain.Factor{
		SourceService:  req.Factor.SourceService,
		ExternalID:     uuid.MustParse(req.Factor.ExternalId),
		BookingID:      uuid.MustParse(req.Factor.BookingId),
		Amount:         uint(req.Factor.TotalAmount),
		Status:         factorDomain.FactorStatusPending,
		Details:        req.Factor.Details,
		InstantPayment: req.Factor.InstantPayment,
		CustomerID:     customerID,
		IsPaid:         false,
	}

	// Validate factor amount
	if factor.Amount <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "total_amount must be greater than zero")
	}

	// Fetch customer's wallet
	customerWallets, err := h.walletService.GetWallets(ctx, domain.WalletFilters{OwnerID: customerID})
	if err != nil || len(customerWallets) == 0 {
		return nil, status.Errorf(codes.NotFound, "customer wallet not found")
	}
	customerWallet := customerWallets[0]

	// Calculate effective balance
	effectiveBalance, err := h.transactionService.GetTransactionSum(ctx, transactionDomain.TransactionFilters{
		WalletID: customerWallet.ID,
		Status:   transactionDomain.TransactionStatusPending,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to calculate effective wallet balance: %v", err)
	}
	totalEffectiveBalance := uint(effectiveBalance) + customerWallet.Balance

	// Add 5% surcharge
	surcharge := factor.Amount / 20
	totalRequired := factor.Amount + surcharge

	if totalEffectiveBalance < totalRequired {
		return nil, status.Errorf(codes.FailedPrecondition, "insufficient balance in customer wallet")
	}

	// Create the factor
	factorID, err := h.factorService.CreateFactor(ctx, factor)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create factor: %v", err)
	}

	// Process distributions
	for _, dist := range req.Factor.Distributions {
		walletID := uuid.MustParse(dist.WalletId)

		transaction := transactionDomain.Transaction{
			WalletID: walletID,
			FactorID: factorID,
			Amount:   uint(dist.Amount),
			Type:     transactionDomain.TransactionTypeCredit,
			Status:   transactionDomain.TransactionStatusPending,
		}
		// TODO: make repo accept batch
		_, err = h.transactionService.CreateTransaction(ctx, transaction)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create transaction for wallet %s: %v", dist.WalletId, err)
		}
		transaction = transactionDomain.Transaction{
			WalletID: customerWallets[0].ID,
			FactorID: factorID,
			Amount:   uint(dist.Amount),
			Type:     transactionDomain.TransactionTypeDebit,
			Status:   transactionDomain.TransactionStatusPending,
		}

		_, err = h.transactionService.CreateTransaction(ctx, transaction)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create transaction for wallet %s: %v", dist.WalletId, err)
		}
	}

	// Create transactions for 5% surcharge
	centralWalletID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	transactions := []transactionDomain.Transaction{
		{
			WalletID: customerWallet.ID,
			FactorID: factorID,
			Amount:   surcharge,
			Type:     transactionDomain.TransactionTypeDebit,
			Status:   transactionDomain.TransactionStatusPending,
		},
		{
			WalletID: centralWalletID,
			FactorID: factorID,
			Amount:   surcharge,
			Type:     transactionDomain.TransactionTypeCredit,
			Status:   transactionDomain.TransactionStatusPending,
		},
	}

	for _, tx := range transactions {
		_, err = h.transactionService.CreateTransaction(ctx, tx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create surcharge transaction: %v", err)
		}
	}

	// Mark factor as paid for instant payment
	if factor.InstantPayment {
		if err := h.factorService.UpdateFactorStatus(ctx, factorID, factorDomain.FactorStatusApproved); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update factor status: %v", err)
		}
	}

	// Return success response
	return &pb.CreateFactorResponse{
		Status:  pb.ResponseStatus_SUCCESS,
		Message: "Factor created successfully.",
		Factor: &pb.Factor{
			Id:             factorID.String(),
			SourceService:  factor.SourceService,
			ExternalId:     factor.ExternalID.String(),
			BookingId:      factor.BookingID.String(),
			TotalAmount:    uint64(factor.Amount),
			CustomerId:     factor.CustomerID.String(),
			Details:        factor.Details,
			InstantPayment: factor.InstantPayment,
			IsPaid:         factor.IsPaid,
			CreatedAt:      factor.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      factor.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (h *GRPCBankHandler) ApplyFactor(ctx context.Context, req *pb.ApplyFactorRequest) (*pb.ApplyFactorResponse, error) {
	// Validate input
	if req.FactorId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "factor ID is required")
	}

	// Parse Factor ID
	factorID, err := uuid.Parse(req.FactorId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid factor ID format: %v", err)
	}

	// Fetch the factor
	factor, err := h.factorService.GetFactorByID(ctx, factorID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve factor: %v", err)
	}

	// Check if the factor is already paid
	if factor.IsPaid {
		return &pb.ApplyFactorResponse{
			Status:  pb.ResponseStatus_FAILED,
			Message: "The factor has already been applied",
		}, nil
	}

	// Fetch the customer's wallet
	customerWallets, err := h.walletService.GetWallets(ctx, domain.WalletFilters{OwnerID: uuid.UUID(factor.CustomerID)})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve customer's wallet: %v", err)
	}
	if len(customerWallets) == 0 {
		return nil, status.Errorf(codes.NotFound, "customer's wallet not found")
	}
	customerWallet := customerWallets[0]

	// Calculate effective balance
	effectiveBalance, err := h.transactionService.GetTransactionSum(ctx, transactionDomain.TransactionFilters{
		WalletID: customerWallet.ID,
		Status:   transactionDomain.TransactionStatusPending,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to calculate effective wallet balance: %v", err)
	}

	if effectiveBalance < 0 {
		return nil, status.Errorf(codes.Internal, "effective balance calculation resulted in a negative value")
	}
	totalEffectiveBalance := uint(effectiveBalance) + customerWallet.Balance

	// Calculate 5% surcharge
	surcharge := factor.Amount / 20
	totalRequired := factor.Amount + surcharge

	// Ensure sufficient balance
	if totalEffectiveBalance < totalRequired {
		return &pb.ApplyFactorResponse{
			Status:  pb.ResponseStatus_FAILED,
			Message: "Insufficient wallet balance",
		}, nil
	}

	// Update the status of all associated pending transactions
	transactions, err := h.transactionService.GetTransactions(ctx, transactionDomain.TransactionFilters{
		FactorID: factor.ID,
		Status:   transactionDomain.TransactionStatusPending,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve pending transactions: %v", err)
	}

	for _, transaction := range transactions {
		err := h.transactionService.UpdateTransactionStatus(ctx, transaction.ID, transactionDomain.TransactionStatusCompleted)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update transaction status: %v", err)
		}
	}

	// Mark the factor as paid
	err = h.factorService.UpdateFactorStatus(ctx, factor.ID, factorDomain.FactorStatusApproved)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update factor status: %v", err)
	}

	// Return success response
	return &pb.ApplyFactorResponse{
		Status:  pb.ResponseStatus_SUCCESS,
		Message: "Factor applied successfully",
	}, nil
}

// GetFactors handles the GetFactors gRPC method.
func (h *GRPCBankHandler) GetFactors(ctx context.Context, req *pb.GetFactorsRequest) (*pb.GetFactorsResponse, error) {
	// Parse FactorId from string to uuid.UUID
	var factorID uuid.UUID
	if req.FactorId != "" {
		parsedFactorID, err := uuid.Parse(req.FactorId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid factor ID format: %v", err)
		}
		factorID = parsedFactorID
	}

	// Convert request filters to domain filters
	filters := factorDomain.FactorFilters{
		FactorID:  factorID,
		BookingID: req.BookingId,
		IsPaid:    &req.IsPaid,
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
	}

	// Call factor service to fetch factors
	factors, totalCount, err := h.factorService.GetFactors(ctx, filters)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve factors: %v", err)
	}

	// Convert domain factors to protobuf Factor messages
	var factorResponses []*pb.Factor
	for _, f := range factors {
		factorResponses = append(factorResponses, &pb.Factor{
			Id:             f.ID.String(),
			SourceService:  f.SourceService,
			ExternalId:     f.ExternalID.String(),
			BookingId:      f.BookingID.String(),
			TotalAmount:    uint64(f.Amount),
			Details:        f.Details,
			InstantPayment: f.InstantPayment,
			IsPaid:         f.IsPaid,
			CreatedAt:      f.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      f.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetFactorsResponse{
		Status:     pb.ResponseStatus_SUCCESS,
		Factors:    factorResponses,
		TotalCount: uint32(totalCount),
		Message:    "Factors retrieved successfully",
	}, nil
}
