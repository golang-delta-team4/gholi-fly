package role

import (
	"context"
	"errors"
	"fmt"
	permissionPort "user-service/internal/permission/port"
	"user-service/internal/role/domain"
	rolePort "user-service/internal/role/port"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/storage/mapper"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var (
	ErrRoleNameNotUnique = errors.New("role name must be unique")
)

type ErrRoleNotFound struct {
	uuid uuid.UUID
}

func (err *ErrRoleNotFound) Error() string {
	return fmt.Sprintf("role %v not found", err.uuid)
}

type service struct {
	repo              rolePort.Repo
	permissionService permissionPort.Service
	userService       userPort.Service
}

func NewService(repo rolePort.Repo, permissionService permissionPort.Service, userService userPort.Service) rolePort.Service {
	return &service{repo: repo, permissionService: permissionService, userService: userService}
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

func (s *service) AssignRole(ctx context.Context, userUUID uuid.UUID, roles []domain.Role) error {
	userID, err := s.userService.GetUserIDByUUID(ctx, userUUID)
	if err != nil {
		return err
	}
	for _, roleDomain := range roles {
		role, err := s.repo.GetRole(ctx, roleDomain.UUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &ErrRoleNotFound{uuid: roleDomain.UUID}
			}
			return err
		}
		err = s.repo.AssignRole(ctx, types.UserRole{UserID: userID, RoleID: role.ID})
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok { //gorm does not have any sentinel error for this error type
				if pgErr.Code == "23505" {
					continue
				}
			}
			return err
		}
	}
	return nil
	
}
