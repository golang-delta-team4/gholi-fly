package grpc

import (
	"context"
	"user-service/api/pb"
	"user-service/api/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcUserHandler struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewGRPCUserHandler(ctx context.Context, userService *service.UserService) *grpcUserHandler {
	return &grpcUserHandler{
		userService: userService,
	}
}

func (h *grpcUserHandler) UserAuthorization(ctx context.Context, req *pb.UserAuthorizationRequest) (*pb.UserAuthorizationResponse, error) {
	ok, err := h.userService.AuthorizeUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user authorization: %v", err)
	}
	if !ok {
		return &pb.UserAuthorizationResponse{AuthorizationStatus: pb.Status_FAILED}, nil
	}
	return &pb.UserAuthorizationResponse{AuthorizationStatus: pb.Status_SUCCESS}, nil
}
