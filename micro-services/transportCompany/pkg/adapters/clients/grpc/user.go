package grpc

import (
	"context"
	"fmt"

	userPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/user"
	clientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcUserClient struct {
	port int
	host string
}

func NewGRPCUserClient(host string, port int) clientPort.GRPCUserClient {
	return &grpcUserClient{host: host, port: port}
}

func (g *grpcUserClient) CheckUserAuthorization(req *userPB.UserAuthorizationRequest) (*userPB.UserAuthorizationResponse, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := userPB.NewUserServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.UserAuthorization(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *grpcUserClient) GetBlockedUser(req *userPB.Empty) (*userPB.GetBlockedUsersResponse, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := userPB.NewUserServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.GetBlockedUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
