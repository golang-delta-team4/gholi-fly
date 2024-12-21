package permission

import (
	"context"
	"errors"
	"user-service/internal/permission/domain"
	permissionPort "user-service/internal/permission/port"
	"user-service/pkg/adapters/storage/mapper"

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
	err = ps.repo.Create(ctx, *permissionType)
	return uuid, err
}
