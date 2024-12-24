package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/google/uuid"
)

type Repo interface {
	CreateTrip(ctx context.Context, trip domain.Trip) (uuid.UUID, error)
	GetTripById(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
	// UpdateTrip(ctx context.Context, trip domain.Trip) error
	// DeleteTrip(ctx context.Context, id uuid.UUID) error
}
