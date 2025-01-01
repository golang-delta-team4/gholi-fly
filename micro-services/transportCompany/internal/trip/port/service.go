package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/google/uuid"
)

type Service interface {
	CreateTrip(ctx context.Context, trip domain.Trip) (uuid.UUID, error)
	GetTripById(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
	GetTrips(ctx context.Context, pageSize int, page int) ([]domain.Trip, error)
	UpdateTrip(ctx context.Context, trip domain.Trip, oldTrip domain.Trip) error
	DeleteTrip(ctx context.Context, id uuid.UUID) error
	ConfirmTrip(ctx context.Context, id uuid.UUID, userId uuid.UUID) error
}
