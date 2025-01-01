package grpc

import (
	"context"
	"fmt"
	notificationClientPort "gholi-fly-hotel/pkg/adapters/clients/grpc/port"

	notificationPB "gholi-fly-hotel/pkg/adapters/clients/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCNotificationClient struct {
	host string
	port int
}

func NewGRPCNotificationClient(host string, port int) notificationClientPort.GRPCNotificationClient {
	return &GRPCNotificationClient{host: host, port: port}
}

func (g *GRPCNotificationClient) AddNotification(req *notificationPB.AddNotificationRequest) (*notificationPB.AddNotificationResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", g.host, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := notificationPB.NewNotificationServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Call the GetUserByToken method
	response, err := client.AddNotification(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
