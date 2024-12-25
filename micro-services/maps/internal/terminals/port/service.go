package port

import (
	"context"
	"gholi-fly-maps/internal/terminals/domain"

	"github.com/google/uuid"
)

// TerminalService defines the interface for terminal business logic.
type TerminalService interface {
	GetAllTerminals(ctx context.Context) ([]domain.Terminal, error)
	CreateTerminal(ctx context.Context, terminal *domain.Terminal) error
	GetTerminalByID(ctx context.Context, id uuid.UUID) (*domain.Terminal, error)
	UpdateTerminal(ctx context.Context, terminal *domain.Terminal) error
	DeleteTerminal(ctx context.Context, id uuid.UUID) error
	SearchTerminals(ctx context.Context, filter TerminalFilter) ([]domain.Terminal, error)
}
