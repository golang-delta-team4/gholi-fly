package role

import (
	"context"
	"errors"
	permissionPort "user-service/internal/permission/port"
	"user-service/internal/role/domain"
	rolePort "user-service/internal/role/port"
	"user-service/pkg/adapters/storage/mapper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrRoleNameNotUnique = errors.New("role name must be unique")
)

type service struct {
	repo rolePort.Repo
	permissionService permissionPort.Service
}

func NewService(repo rolePort.Repo, permissionService permissionPort.Service) rolePort.Service {
	return &service{repo: repo, permissionService: permissionService}
}

func (s *service) CreateRole(ctx context.Context, role *domain.Role) (uuid.UUID, error) {
	roleType := mapper.RoleDomain2Storage(*role)
	permissions, err := s.permissionService.GetPermissionsByUUID(ctx, role.Permissions)
	if err != nil {
		return uuid.Nil, err
	}
	NewUUID := uuid.New()
	roleType.UUID = NewUUID
	roleType.Permissions = permissions
	_, err = s.repo.Create(ctx, roleType)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok { //gorm does not have any sentinel error for this error type
			if pgErr.Code == "23505" {
				return uuid.Nil, ErrRoleNameNotUnique
			}
		}
		return uuid.Nil, err
	}
	return NewUUID, nil
}