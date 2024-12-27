package grpc

import (
	"context"
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/pb"
	bankClientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"

	"google.golang.org/grpc"
)

type GRPCBankClient struct {
	port int
	host string
}

func NewGRPCBankClient(host string, port int) bankClientPort.GRPCBankClient {
	return &GRPCBankClient{host: host, port: port}
}

func (g *GRPCBankClient) CreateUserWallet(userUUID string) (*pb.CreateWalletResponse, error) {

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := pb.NewWalletServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Prepare the request
	request := &pb.CreateWalletRequest{
		OwnerId: userUUID,
		Type:    pb.WalletType_PERSON,
	}

	// Call the GetUserByToken method
	response, err := client.CreateWallet(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GRPCBankClient) CreateFactor(req *pb.CreateFactorRequest) (*pb.CreateFactorResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := pb.NewFactorServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.CreateFactor(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
