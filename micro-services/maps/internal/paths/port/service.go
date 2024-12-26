package port

import (
	"context"
	"gholi-fly-maps/internal/paths/domain"

	"github.com/google/uuid"
)

type PathService interface {
	GetAllPaths(ctx context.Context) ([]domain.Path, error)
	CreatePath(ctx context.Context, path *domain.Path) error
	UpdatePath(ctx context.Context, path *domain.Path) error
	DeletePath(ctx context.Context, id uuid.UUID) error
}
