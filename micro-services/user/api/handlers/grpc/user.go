package grpc

import (
	"context"
	"user-service/api/pb"
	"user-service/api/presenter"
	"user-service/api/service"

	"github.com/google/uuid"
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
	userUUID, err := uuid.Parse(req.UserUUID)
	ok, err := h.userService.AuthorizeUser(ctx, presenter.UserAuthorization{UserUUID: userUUID, Route: req.Route, Method: req.Method})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user authorization: %v", err)
	}
	if !ok {
		return &pb.UserAuthorizationResponse{AuthorizationStatus: pb.Status_FAILED}, nil
	}
	return &pb.UserAuthorizationResponse{AuthorizationStatus: pb.Status_SUCCESS}, nil
}

func (h *grpcUserHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user authorization: %v", err)
	}
	return &pb.GetUserResponse{Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, Uuid: user.UUID.String()}, nil
}

func (h *grpcUserHandler) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUserByUUID(ctx, req.UserUUID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user authorization: %v", err)
	}
	return &pb.GetUserResponse{Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, Uuid: user.UUID.String()}, nil
}

func (h *grpcUserHandler) GetBlockedUsers(ctx context.Context, req *pb.Empty) (*pb.GetBlockedUsersResponse, error) {
	uuids, err := h.userService.GetBlockedUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &pb.GetBlockedUsersResponse{Uuids: uuids}, nil
}
