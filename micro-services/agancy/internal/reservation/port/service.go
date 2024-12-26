package port

import (
	"context"

	"gholi-fly-agancy/internal/reservation/domain"

	"github.com/google/uuid"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, reservation domain.Reservation) (domain.ReservationID, error)
	GetReservationByID(ctx context.Context, id domain.ReservationID) (*domain.Reservation, error)
	UpdateReservation(ctx context.Context, reservation domain.Reservation) error
	DeleteReservation(ctx context.Context, id domain.ReservationID) error
	ListReservationsBytour(ctx context.Context, tourID uuid.UUID) ([]domain.Reservation, error)
	ListReservationsByUser(ctx context.Context, userID uuid.UUID) ([]domain.Reservation, error)
}
