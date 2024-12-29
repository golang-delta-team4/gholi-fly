package port

import (
	"context"
	"vehicle/internal/vehicle/domain"

	"github.com/google/uuid"
)

// VehicleService defines the interface for the service layer.
type VehicleService interface {
	CreateVehicle(ctx context.Context, vehicle *domain.Vehicle) error
	GetVehicleByID(ctx context.Context, id uuid.UUID) (*domain.Vehicle, error)
	GetAllVehicles(ctx context.Context) ([]domain.Vehicle, error)
	UpdateVehicle(ctx context.Context, vehicle *domain.Vehicle) error
	DeleteVehicle(ctx context.Context, id uuid.UUID) error
	ProcessTripRequest(ctx context.Context) (*domain.TripRequest, error) 
	MatchVehicle(ctx context.Context, tripRequest *domain.TripRequest) (*domain.Vehicle, error)
}
