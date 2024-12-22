package grpc

import (
	"context"
	"fmt"
	"gholi-fly-bank/api/handlers/grpc/handlers"
	"gholi-fly-bank/api/handlers/grpc/middlewares"
	"gholi-fly-bank/api/pb"
	"gholi-fly-bank/app"
	"gholi-fly-bank/config"

	"net"

	"google.golang.org/grpc"
)

// Run initializes and starts the gRPC server.
func Run(app app.App, cfg config.Config) error {
	methodsRequiringTx := map[string]bool{
		"/wallet.WalletService/CreateWallet": true,
		"/wallet.WalletService/GetWallets":   false,
	}
	// Create gRPC server with interceptors
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middlewares.LoggerInterceptor(),                                  // Log requests and responses
			middlewares.TransactionInterceptor(app.DB(), methodsRequiringTx), // Attach DB transactions to context
		),
	)

	// Register gRPC handlers
	pb.RegisterWalletServiceServer(server, handlers.NewGRPCBankHandler(context.Background(), app))

	// Start listening
	address := fmt.Sprintf(":%d", cfg.Server.GrpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}

	fmt.Printf("gRPC server is running on %s\n", address)

	// Serve gRPC
	return server.Serve(listener)
}
