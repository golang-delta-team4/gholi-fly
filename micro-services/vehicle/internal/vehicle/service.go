package vehicle

import (
	"context"
	"fmt"
	"vehicle/internal/vehicle/domain"
	"vehicle/internal/vehicle/port"

	"github.com/google/uuid"
)

type service struct {
	repo port.VehicleRepository
}

// NewVehicleService creates a new instance of the VehicleService.
func NewVehicleService(repo port.VehicleRepository) port.VehicleService {
	return &service{repo: repo}
}

// Implement the methods defined in the VehicleService interface.
func (s *service) CreateVehicle(ctx context.Context, vehicle *domain.Vehicle) error {
	if vehicle.UniqueCode == "" {
		return fmt.Errorf("unique code is required")
	}
	return s.repo.Create(ctx, vehicle)
}

func (s *service) GetVehicleByID(ctx context.Context, id uuid.UUID) (*domain.Vehicle, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) GetAllVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) UpdateVehicle(ctx context.Context, vehicle *domain.Vehicle) error {
	return s.repo.Update(ctx, vehicle)
}

func (s *service) DeleteVehicle(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
