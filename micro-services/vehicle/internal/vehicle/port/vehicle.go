package port

import (
	"context"
	"vehicle/internal/vehicle/domain"
	"vehicle/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

type VehicleRepository interface {
    Create(ctx context.Context, vehicle *domain.Vehicle) error
    GetByID(ctx context.Context, id uuid.UUID) (*domain.Vehicle, error)
    GetAll(ctx context.Context) ([]domain.Vehicle, error)
    GetMatchedVehicle(ctx context.Context, vehicleMatchRequest *domain.MatchMakerRequest) (types.Vehicle, error)
    Update(ctx context.Context, vehicle *domain.Vehicle) error
    Delete(ctx context.Context, id uuid.UUID) error
    ProcessTripRequest(ctx context.Context) (*domain.TripRequest, error) // Updated to match
    CreateReservation(ctx context.Context, vehicleReserve types.VehicleReserve) (uuid.UUID, error)
}



