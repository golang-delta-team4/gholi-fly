package grpc

import (
	"context"
	"fmt"

	rolePB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/role"
	roleClientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCRoleClient struct {
	port int
	host string
}

func NewGRPCRoleClient(host string, port int) roleClientPort.GRPCRoleClient {
	return &GRPCRoleClient{host: host, port: port}
}

func (g *GRPCRoleClient) CreateRole(req *rolePB.GrantResourceAccessRequest) (*rolePB.GrantResourceAccessResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := rolePB.NewRoleServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.GrantResourceAccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
