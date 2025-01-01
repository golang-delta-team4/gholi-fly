package service

import (
	"context"
	"strings"
	"user-service/api/presenter"
	"user-service/internal/permission/domain"
	permissionPort "user-service/internal/permission/port"

	"github.com/google/uuid"
)

type PermissionService struct {
	service permissionPort.Service
}

func NewPermissionService(service permissionPort.Service) *PermissionService {
	return &PermissionService{service: service}
}

func (ps *PermissionService) Create(ctx context.Context, permission *presenter.CreatePermissionRequest) (uuid.UUID, error) {
	return ps.service.CreatePermission(ctx, &domain.Permission{Route: strings.ToLower(permission.Route), Method: strings.ToLower(permission.Method)})
}
