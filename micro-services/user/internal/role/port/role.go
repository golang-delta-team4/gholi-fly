package port

import (
	"context"
	"user-service/pkg/adapters/storage/types"
)

type Repo interface {
	Create(ctx context.Context,role *types.Role) (uint, error)
}