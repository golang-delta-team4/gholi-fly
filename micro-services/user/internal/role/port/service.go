package port

import (
	"context"
	"user-service/internal/role/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateRole(ctx context.Context,role *domain.Role) (uuid.UUID, error)
}