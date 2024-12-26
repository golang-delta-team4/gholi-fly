package reservation

import (
	"context"
	"errors"
	"fmt"

	"gholi-fly-agancy/internal/reservation/domain"
	"gholi-fly-agancy/internal/reservation/port"

	"github.com/google/uuid"
)

var (
	ErrReservationCreateFailed     = errors.New("failed to create reservation")
	ErrReservationNotFound         = errors.New("reservation not found")
	ErrReservationValidationFailed = errors.New("reservation validation failed")
)

type service struct {
	repo port.ReservationRepo
}

func NewService(repo port.ReservationRepo) port.ReservationService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateReservation(ctx context.Context, reservation domain.Reservation) (domain.ReservationID, error) {
	if err := reservation.Validate(); err != nil {
		return domain.ReservationID{}, fmt.Errorf("%w: %v", ErrReservationValidationFailed, err)
	}

	reservationID, err := s.repo.Create(ctx, reservation)
	if err != nil {
		return domain.ReservationID{}, fmt.Errorf("%w: %v", ErrReservationCreateFailed, err)
	}

	return reservationID, nil
}

func (s *service) GetReservationByID(ctx context.Context, id domain.ReservationID) (*domain.Reservation, error) {
	reservation, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if reservation == nil {
		return nil, ErrReservationNotFound
	}

	return reservation, nil
}

func (s *service) UpdateReservation(ctx context.Context, reservation domain.Reservation) error {
	if err := reservation.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrReservationValidationFailed, err)
	}

	if err := s.repo.Update(ctx, reservation); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteReservation(ctx context.Context, id domain.ReservationID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ListReservationsBytour(ctx context.Context, tourID uuid.UUID) ([]domain.Reservation, error) {
	reservations, err := s.repo.ListBytourID(ctx, tourID)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (s *service) ListReservationsByUser(ctx context.Context, userID uuid.UUID) ([]domain.Reservation, error) {
	reservations, err := s.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}
