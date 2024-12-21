package port

import (
	"context"
	"user-service/pkg/adapters/storage/types"
)

type Repo interface {
	Create(ctx context.Context, permission types.Permission) error
	CheckPermissionExistence(ctx context.Context, route string, method string) (bool, error)
}
