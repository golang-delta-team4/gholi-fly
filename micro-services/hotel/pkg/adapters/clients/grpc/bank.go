package grpc

import (
	"context"
	"fmt"
	bankClientPort "gholi-fly-hotel/pkg/adapters/clients/grpc/port"

	bankPB "gholi-fly-hotel/pkg/adapters/clients/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCBankClient struct {
	host string
	port int
}

func NewGRPCBankClient(host string, port int) bankClientPort.GRPCBankClient {
	return &GRPCBankClient{host: host, port: port}
}

func (g *GRPCBankClient) GetUserWallets(req *bankPB.GetWalletsRequest) (*bankPB.GetWalletsResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := bankPB.NewWalletServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.GetWallets(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GRPCBankClient) CreateFactor(req *bankPB.CreateFactorRequest) (*bankPB.CreateFactorResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := bankPB.NewFactorServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.CreateFactor(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GRPCBankClient) ApplyFactor(req *bankPB.ApplyFactorRequest) (*bankPB.ApplyFactorResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := bankPB.NewFactorServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.ApplyFactor(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
