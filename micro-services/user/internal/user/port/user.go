package port

import (
	"context"
	"user-service/pkg/adapters/storage/types"
)

type Repo interface {
	Create(ctx context.Context, user types.User) error
	GetByEmail(ctx context.Context, email string) (*types.User, error)
	UpdateRefreshToken(ctx context.Context, userID uint, refreshToken string) error
}
