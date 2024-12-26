package grpc

import (
	"context"
	"fmt"
	"user-service/api/pb"
	"user-service/api/service"
	"user-service/app"
	"user-service/config"

	"net"

	"google.golang.org/grpc"
)

// Run initializes and starts the gRPC server.
func Run(app app.App, cfg config.Config) error {
	userService := service.NewUserService(app.UserService(), cfg.Server.AuthExpMinute, cfg.Server.AuthRefreshMinute, cfg.Server.Secret)
	// Create gRPC server with interceptors
	server := grpc.NewServer()

	// Register gRPC handlers
	pb.RegisterUserServiceServer(server, NewGRPCUserHandler(context.Background(), userService))
	// Start listening
	address := fmt.Sprintf(":%d", cfg.Server.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}

	fmt.Printf("gRPC server is running on %s\n", address)

	// Serve gRPC
	return server.Serve(listener)
}
