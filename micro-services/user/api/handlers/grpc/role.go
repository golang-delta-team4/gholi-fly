package grpc

import (
	"context"
	rolePB "user-service/api/pb/role"
	"user-service/api/service"
	permissionDomain "user-service/internal/permission/domain"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcRoleHandler struct {
	rolePB.UnimplementedRoleServiceServer
	roleService *service.RoleService
}

func NewGRPCRoleHandler(ctx context.Context, roleService *service.RoleService) *grpcRoleHandler {
	return &grpcRoleHandler{
		roleService: roleService,
	}
}

func (h *grpcRoleHandler) GrantResourceAccess(ctx context.Context, req *rolePB.GrantResourceAccessRequest) (*rolePB.GrantResourceAccessResponse, error) {
	ownerUUID, err := uuid.Parse(req.OwnerUUID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse owner uuid")
	}
	var permissions []permissionDomain.Permission
	for _, permission := range req.Permissions {
		permissions = append(permissions, permissionDomain.Permission{Route: permission.Route, Method: permission.Method})
	}
	err = h.roleService.GrantResourceAccess(ctx, ownerUUID, permissions, req.RoleName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user authorization: %v", err)
	}
	return &rolePB.GrantResourceAccessResponse{Status: rolePB.AccessStatus_ACCESS_GRANTED}, nil
}
