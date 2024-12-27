package permission

import (
	"context"
	"errors"
	"user-service/internal/permission/domain"
	permissionPort "user-service/internal/permission/port"
	"user-service/pkg/adapters/storage/mapper"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

type service struct {
	repo permissionPort.Repo
}

func NewService(repo permissionPort.Repo) permissionPort.Service {
	return &service{repo: repo}
}

var (
	ErrAlreadyExists = errors.New("permission already exists")
)

func (ps *service) CreatePermission(ctx context.Context, permission *domain.Permission) (uuid.UUID, error) {
	exist, err := ps.repo.CheckPermissionExistence(ctx, permission.Route, permission.Method)
	if err != nil {
		return uuid.Nil, err
	}
	if exist {
		return uuid.Nil, ErrAlreadyExists
	}
	uuid := uuid.New()
	permissionType := mapper.PermissionDomain2Storage(*permission)
	permissionType.UUID = uuid
	err = ps.repo.Create(ctx, []types.Permission{*permissionType})
	return uuid, err
}

func (ps *service) GetPermissionsByUUID(ctx context.Context, permissions []domain.Permission) ([]types.Permission, error) {
	var typedPermission []types.Permission
	for _, permission := range permissions {
		typedPermission = append(typedPermission, types.Permission{UUID: permission.UUID})
	}
	return ps.repo.GetPermissionsByUUID(ctx, typedPermission)
}

func (ps *service) CreatePermissions(ctx context.Context, permissions []domain.Permission) ([]domain.Permission, error) {
	var permissionTypes []types.Permission
	// var permissionsWithUUID []domain.Permission
	for _, permission := range permissions {
		uuid := uuid.New()
		permission.UUID = uuid
		permissionType := mapper.PermissionDomain2Storage(permission)
		permissionTypes = append(permissionTypes, *permissionType)
	}
	err := ps.repo.Create(ctx, permissionTypes)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (ps *service) GetAllPermissions(ctx context.Context) ([]types.Permission, error) {
	return ps.repo.GetAllPermissions(ctx)
}
