package port

import (
	"context"
	"vehicle/internal/vehicle/domain"

	"github.com/google/uuid"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle *domain.Vehicle) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Vehicle, error)
	GetAll(ctx context.Context) ([]domain.Vehicle, error)
	Update(ctx context.Context, vehicle *domain.Vehicle) error
	Delete(ctx context.Context, id uuid.UUID) error
}


