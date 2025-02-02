package role

import (
	"context"
	"errors"
	"fmt"
	"user-service/api/presenter"
	permissionDomain "user-service/internal/permission/domain"
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

func (s *service) GrantResourceAccess(ctx context.Context, ownerUUID uuid.UUID, permissions []permissionDomain.Permission, roleName string) error {
	permissionsWithUUID, err := s.permissionService.CreatePermissions(ctx, permissions)
	if err != nil {
		return err
	}
	roleUUID, err := s.CreateRole(ctx, &domain.Role{Name: roleName, Permissions: permissionsWithUUID})
	if err != nil {
		return err
	}
	err = s.AssignRole(ctx, ownerUUID, []domain.Role{{UUID: roleUUID}})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateSuperAdminRole(ctx context.Context) (uuid.UUID, error) {
	permissions, err := s.permissionService.GetAllPermissions(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	role, err := s.repo.GetByName(ctx, "SuperAdmin")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		roleUUID := uuid.New()
		_, err = s.repo.Create(ctx, &types.Role{Name: "SuperAdmin", UUID: roleUUID, Permissions: permissions})
		if err != nil {
			return uuid.Nil, err
		}
		return roleUUID, nil
	}
	if err != nil {
		return uuid.Nil, err
	}
	return role.UUID, nil
}

func (us *service) GetAllRoles(ctx context.Context, query presenter.PaginationQuery) ([]domain.Role, error) {
	roles, err := us.repo.GetAllRoles(ctx, query)
	if err != nil {
		return nil, err
	}
	storageRoles := mapper.RoleStorageList2DomainList(roles)
	return storageRoles, nil
}

func (us *service) DeleteRole(ctx context.Context, roleUUID uuid.UUID) error {
	role, err := us.repo.GetRole(ctx, roleUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ErrRoleNotFound{roleUUID}
		}
	}
	return us.repo.Delete(ctx, *role)
}
