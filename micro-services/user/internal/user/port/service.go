package port

import (
	"context"
	"time"
	"user-service/api/presenter"
	"user-service/internal/user/domain"

	"github.com/google/uuid"
)

type Service interface {
	SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error)
	SignIn(ctx context.Context, user *domain.User) (uuid.UUID, error)
	UpdateUserRefreshToken(ctx context.Context, userID uint, refreshToken string, until time.Time) error
	GetUserRefreshToken(ctx context.Context, userID uint) (string, error)
	ValidateRefreshToken(ctx context.Context, userID uint, refreshToken string) (bool, error)
	GetUserIDByUUID(ctx context.Context, userUUID uuid.UUID) (uint, error)
	GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*domain.User, error)
	BlockUser(ctx context.Context, userUUID uuid.UUID) (error)
	UnBlockUser(ctx context.Context, userUUID uuid.UUID) (error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	AuthorizeUser(ctx context.Context, userAuthorization *domain.UserAuthorize) (bool, error)
	GetAllUsers(ctx context.Context, query presenter.PaginationQuery) ([]domain.User, error)
}
