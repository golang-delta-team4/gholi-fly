package paths

import (
	"context"
	"fmt"
	"gholi-fly-maps/internal/paths/domain"
	port_p "gholi-fly-maps/internal/paths/port"
	port_t "gholi-fly-maps/internal/terminals/port"

	"github.com/google/uuid"
)

type service struct {
	repo         port_p.PathRepository
	terminalRepo port_t.TerminalRepository // Add terminalRepo to the struct
}

func NewPathService(repo port_p.PathRepository, terminalRepo port_t.TerminalRepository) port_p.PathService {
	return &service{
		repo:         repo,
		terminalRepo: terminalRepo,
	}
}

func (s *service) GetAllPaths(ctx context.Context) ([]domain.Path, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) CreatePath(ctx context.Context, path *domain.Path) (*domain.Path, error) {
	if path.ID == uuid.Nil {
		path.ID = uuid.New()
	}
	return s.repo.Create(ctx, path)
}

func (s *service) UpdatePath(ctx context.Context, path *domain.Path) error {
	// Validate that the path exists
	existingPath, err := s.repo.GetByID(ctx, path.ID.String())
	if err != nil || existingPath == nil {
		return fmt.Errorf("path not found")
	}

	// Validate terminal IDs (if provided)
	if path.SourceTerminalID != uuid.Nil {
		if _, err := s.terminalRepo.GetByID(ctx, path.SourceTerminalID); err != nil {
			return fmt.Errorf("source terminal does not exist")
		}
	}
	if path.DestinationTerminalID != uuid.Nil {
		if _, err := s.terminalRepo.GetByID(ctx, path.DestinationTerminalID); err != nil {
			return fmt.Errorf("destination terminal does not exist")
		}
	}

	// Validate route code uniqueness (if provided)
	if path.RouteCode != "" && path.RouteCode != existingPath.RouteCode {
		paths, err := s.repo.GetAll(ctx)
		if err != nil {
			return err
		}
		for _, p := range paths {
			if p.RouteCode == path.RouteCode {
				return fmt.Errorf("route code must be unique")
			}
		}
	}

	// Validate distance
	if path.DistanceKM < 0 {
		return fmt.Errorf("distance must be a positive number")
	}

	// Delegate to repository
	return s.repo.Update(ctx, path)
}

func (s *service) DeletePath(ctx context.Context, id uuid.UUID) error {
	// Validate that the path exists
	existingPath, err := s.repo.GetByID(ctx, id.String())
	if err != nil || existingPath == nil {
		return fmt.Errorf("path not found")
	}

	// Delegate to the repository to delete the path
	return s.repo.Delete(ctx, id.String())
}

func (s *service) FilterPaths(ctx context.Context, filters map[string]interface{}) ([]domain.Path, error) {
	return s.repo.FilterPaths(ctx, filters)
}
