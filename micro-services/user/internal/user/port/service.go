package port

import (
	"context"
	"user-service/internal/user/domain"
	"github.com/google/uuid"
)


type Service interface {
	SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error)
}