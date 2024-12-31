package port

import (
	"context"
	"user-service/pkg/adapters/storage/types"
)

type Repo interface {
	Create(ctx context.Context, permission []types.Permission) error
	Get(ctx context.Context, permission types.Permission) error
	CheckPermissionExistence(ctx context.Context, route string, method string) (bool, error)
	GetPermissionsByUUID(ctx context.Context, permissions []types.Permission) ([]types.Permission, error)
	GetAllPermissions(ctx context.Context) ([]types.Permission, error)
}
