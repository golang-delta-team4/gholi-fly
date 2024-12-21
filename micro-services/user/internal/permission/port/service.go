package port

import (
	"context"
	"user-service/internal/permission/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreatePermission(ctx context.Context, permission *domain.Permission) (uuid.UUID, error)
}
