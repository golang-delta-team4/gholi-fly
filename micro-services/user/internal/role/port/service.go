package port

import (
	"context"
	"user-service/api/presenter"
	permissionDomain "user-service/internal/permission/domain"
	"user-service/internal/role/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateRole(ctx context.Context, role *domain.Role) (uuid.UUID, error)
	CreateSuperAdminRole(ctx context.Context) (error)
	AssignRole(ctx context.Context, userUUID uuid.UUID, roles []domain.Role) error
	DeleteRole(ctx context.Context, roleUUID uuid.UUID) error
	GrantResourceAccess(ctx context.Context, ownerUUID uuid.UUID, permissions []permissionDomain.Permission, roleName string) error
	GetAllRoles(ctx context.Context, query presenter.PaginationQuery) ([]domain.Role, error)
}
