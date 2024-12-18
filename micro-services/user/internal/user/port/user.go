package port

import (
	"context"
	"user-service/pkg/adapters/storage/types"
)


type Repo interface {
	Create(ctx context.Context, user types.User) error
}