package port

import (
	"context"
	"user-service/api/presenter"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, role *types.Role) (uint, error)
	AssignRole(ctx context.Context, userRole types.UserRole) error
	GetRole(ctx context.Context, roleUUID uuid.UUID) (*types.Role, error)
	GetAllRoles(ctx context.Context, query presenter.PaginationQuery) ([]types.Role, error)
}
