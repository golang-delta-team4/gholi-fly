package port

import (
	"context"
	"gholi-fly-maps/internal/paths/domain"
)

type PathRepository interface {
	GetAll(ctx context.Context) ([]domain.Path, error)
	Create(ctx context.Context, path *domain.Path) error
	GetByID(ctx context.Context, id string) (*domain.Path, error)
	Update(ctx context.Context, path *domain.Path) error
	Delete(ctx context.Context, id string) error
	FilterPaths(ctx context.Context, filters map[string]interface{}) ([]domain.Path, error) // New method
}
