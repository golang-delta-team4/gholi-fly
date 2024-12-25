package terminals

import (
	"context"
	"fmt"
	"gholi-fly-maps/internal/terminals/domain"
	"gholi-fly-maps/internal/terminals/port"

	"github.com/google/uuid"
)

type service struct {
	repo port.TerminalRepository
}

// NewTerminalService creates a new TerminalService instance.
func NewTerminalService(repo port.TerminalRepository) port.TerminalService {
	return &service{repo: repo}
}

func (s *service) GetAllTerminals(ctx context.Context) ([]domain.Terminal, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) CreateTerminal(ctx context.Context, terminal *domain.Terminal) error {
	return s.repo.Create(ctx, terminal)
}

func (s *service) GetTerminalByID(ctx context.Context, id uuid.UUID) (*domain.Terminal, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) UpdateTerminal(ctx context.Context, terminal *domain.Terminal) error {
	// Validation
	if terminal.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if terminal.Location == "" {
		return fmt.Errorf("location cannot be empty")
	}
	validTypes := map[string]bool{
		"bus_station":   true,
		"airport":       true,
		"train_station": true,
		"sea_port":      true,
	}
	if !validTypes[terminal.Type] {
		return fmt.Errorf("invalid terminal type")
	}

	// Delegate to the repository
	return s.repo.Update(ctx, terminal)
}

func (s *service) DeleteTerminal(ctx context.Context, id uuid.UUID) error {
	// Fetch the terminal to ensure it exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("terminal not found")
	}

	// Check user permissions (placeholder)
	// TODO: Add permission check logic
	// if !userHasPermission(ctx, terminal) {
	//     return fmt.Errorf("user does not have permission to delete this terminal")
	// }

	// Proceed with deletion
	return s.repo.Delete(ctx, id)
}

// filter
func (s *service) SearchTerminals(ctx context.Context, filter port.TerminalFilter) ([]domain.Terminal, error) {
	// Validation
	if filter.ID == "" && filter.Name == "" && filter.City == "" && filter.Type == "" {
		return nil, fmt.Errorf("at least one filter parameter must be provided")
	}

	// Delegate to the repository
	return s.repo.Search(ctx, filter)
}

