package port

import (
	"context"
	"gholi-fly-maps/internal/terminals/domain"

	"github.com/google/uuid"
)

// TerminalRepository defines the interface for terminal data access.
type TerminalRepository interface {
	GetAll(ctx context.Context) ([]domain.Terminal, error)
	Create(ctx context.Context, terminal *domain.Terminal) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Terminal, error)
	Update(ctx context.Context, terminal *domain.Terminal) error
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, filter TerminalFilter) ([]domain.Terminal, error)
}

type TerminalFilter struct {
	ID   string
	Name string
	City string
	Type string
}