package port

import (
	"context"

	"gholi-fly-agancy/internal/reservation/domain"

	"github.com/google/uuid"
)

type ReservationRepo interface {
	Create(ctx context.Context, reservation domain.Reservation) (domain.ReservationID, error)
	GetByID(ctx context.Context, id domain.ReservationID) (*domain.Reservation, error)
	Update(ctx context.Context, reservation domain.Reservation) error
	Delete(ctx context.Context, id domain.ReservationID) error
	ListByTourID(ctx context.Context, tourID uuid.UUID) ([]domain.Reservation, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Reservation, error)
}
