package service

import (
	"context"
	"fmt"
	"notification-nats/config"

	userPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetUserByUUID(userUUID string, cfg config.Config) (*userPB.GetUserResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", cfg.UserGRPC.Host, cfg.UserGRPC.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := userPB.NewUserServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Prepare the request
	request := &userPB.GetUserByUUIDRequest{
		UserUUID: userUUID,
	}

	// Call the GetUserByToken method
	response, err := client.GetUserByUUID(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
