package port

import (
	"context"
	"user-service/internal/permission/domain"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

type Service interface {
	CreatePermission(ctx context.Context, permission *domain.Permission) (uuid.UUID, error)
	CreatePermissions(ctx context.Context, permission []domain.Permission) ([]domain.Permission, error)
	GetPermissionsByUUID(ctx context.Context, permissions []domain.Permission) ([]types.Permission, error)
}
