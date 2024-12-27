package port

import (
	"context"
	permissionDomain "user-service/internal/permission/domain"
	"user-service/internal/role/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateRole(ctx context.Context, role *domain.Role) (uuid.UUID, error)
	AssignRole(ctx context.Context, userUUID uuid.UUID, roles []domain.Role) error
	GrantResourceAccess(ctx context.Context, ownerUUID uuid.UUID, permissions []permissionDomain.Permission, roleName string) error
}
