package service

import (
	"context"
	"fmt"
	"user-service/api/presenter"
	permissionDomain "user-service/internal/permission/domain"
	"user-service/internal/role/domain"
	rolePort "user-service/internal/role/port"

	"github.com/google/uuid"
)

type ErrInvalidUUID struct {
	uuid string
}

func (err *ErrInvalidUUID) Error() string {
	return fmt.Sprintf("invalid UUID format %v", err.uuid)
}

type RoleService struct {
	service rolePort.Service
}

func NewRoleService(service rolePort.Service) *RoleService {
	return &RoleService{service: service}
}

func (ps *RoleService) Create(ctx context.Context, role *presenter.CreateRoleRequest) (uuid.UUID, error) {
	var permissions []permissionDomain.Permission
	for _, permissionUUID := range role.PermissionUUIDs {
		parsedUUID, err := uuid.Parse(permissionUUID)
		if err != nil {
			return uuid.Nil, &ErrInvalidUUID{uuid: permissionUUID} // TODO: return list of errors
		}
		permissions = append(permissions, permissionDomain.Permission{UUID: parsedUUID})
	}
	return ps.service.CreateRole(ctx, &domain.Role{Name: role.Name, Permissions: permissions})
}

