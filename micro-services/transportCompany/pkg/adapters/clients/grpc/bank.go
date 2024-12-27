package grpc

import (
	"context"
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/pb"
	bankClientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCBankClient struct {
	port int
	host string
}

func NewGRPCBankClient(host string, port int) bankClientPort.GRPCBankClient {
	return &GRPCBankClient{host: host, port: port}
}

func (g *GRPCBankClient) CreateFactor(req *pb.CreateFactorRequest) (*pb.CreateFactorResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
