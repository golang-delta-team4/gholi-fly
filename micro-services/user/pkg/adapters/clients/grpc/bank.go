package grpc

import (
	"context"
	"fmt"
	bankPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/bank"
	bankClientPort "user-service/pkg/adapters/clients/grpc/port"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCBankClient struct {
	port int
	host   string
}

func NewGRPCBankClient(host string, port int) bankClientPort.GRPCBankClient {
	return &GRPCBankClient{host: host, port: port}
}

func (g *GRPCBankClient) CreateUserWallet(userUUID string) (*bankPB.CreateWalletResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := bankPB.NewWalletServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Prepare the request
	request := &bankPB.CreateWalletRequest{
		OwnerId: userUUID,
		Type:    bankPB.WalletType_PERSON,
	}

	// Call the GetUserByToken method
	response, err := client.CreateWallet(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
