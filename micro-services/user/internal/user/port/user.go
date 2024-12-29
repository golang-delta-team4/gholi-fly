package port

import (
	"context"
	"user-service/api/presenter"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, user types.User) error
	GetByEmail(ctx context.Context, email string) (*types.User, error)
	UpdateRefreshToken(ctx context.Context, refreshToken types.RefreshToken) error
	AddRefreshToken(ctx context.Context, refreshToken *types.RefreshToken) error
	DeleteRefreshToken(ctx context.Context, userID uint) error
	GetRefreshToken(ctx context.Context, userID uint) (types.RefreshToken, error)
	GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*types.User, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	AuthorizeUser(ctx context.Context, userAuthorization *types.UserAuthorization) (bool, error)
	GetAllUsers(ctx context.Context, query presenter.PaginationQuery) ([]types.User, error)
}
