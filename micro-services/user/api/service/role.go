package service

import (
	"context"
	"fmt"
	"user-service/api/presenter"
	permissionDomain "user-service/internal/permission/domain"
	roleDomain "user-service/internal/role/domain"
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
		permissions = append(permissions, permissionDomain.Permission{UUID: permissionUUID})
	}
	return ps.service.CreateRole(ctx, &domain.Role{Name: role.Name, Permissions: permissions})
}

func (ps *RoleService) Assign(ctx context.Context, assignRolePresenter *presenter.AssignRoleRequest) (error) {
	var roles []roleDomain.Role
	for _, roleUUID := range assignRolePresenter.RoleUUIDs {
		roles = append(roles, roleDomain.Role{UUID: roleUUID})
	}
	return ps.service.AssignRole(ctx, assignRolePresenter.UserUUID, roles)
}

