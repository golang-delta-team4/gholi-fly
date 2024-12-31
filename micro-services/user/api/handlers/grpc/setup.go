package grpc

import (
	"context"
	"fmt"
	userPB "user-service/api/pb"
	rolePB "user-service/api/pb/role"
	"user-service/api/service"
	"user-service/app"
	"user-service/config"

	"net"

	"google.golang.org/grpc"
)

// Run initializes and starts the gRPC server.
func Run(app app.App, cfg config.Config) error {
	
	userService := service.NewUserService(app.UserService(context.Background()), cfg.Server.AuthExpMinute, cfg.Server.AuthRefreshMinute, cfg.Server.Secret)
	roleService := service.NewRoleService(app.RoleService(context.Background()))
	// Create gRPC server with interceptors
	server := grpc.NewServer()

	// Register gRPC handlers
	userPB.RegisterUserServiceServer(server, NewGRPCUserHandler(context.Background(), userService))
	rolePB.RegisterRoleServiceServer(server, NewGRPCRoleHandler(context.Background(), roleService))
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
