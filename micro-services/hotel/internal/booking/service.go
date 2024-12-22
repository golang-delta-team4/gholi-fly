package booking

import (
	"context"
	"errors"
	"gholi-fly-hotel/internal/booking/domain"
	"gholi-fly-hotel/internal/booking/port"
)

var (
	ErrBookingCreation      = errors.New("error on creating booking")
	ErrBookingNotFound      = errors.New("booking not found")
	ErrInvalidSourceService = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// CreateBooking creates a new booking
func (s *service) CreateBooking(ctx context.Context, booking domain.Booking) (domain.BookingUUID, error) {
	bookingID, err := s.repo.Create(ctx, booking)
	if err != nil {
		return domain.BookingUUID{}, ErrBookingCreation
	}
	return bookingID, nil
}

// GetBookingByID returns a booking by its ID
func (s *service) GetBookingByID(ctx context.Context, bookingID domain.BookingUUID) (*domain.Booking, error) {
	booking, err := s.repo.GetByID(ctx, bookingID)
	if err != nil {
		return nil, ErrBookingNotFound
	}
	return booking, nil
}

// GetBookings returns all bookings
func (s *service) GetBookings(ctx context.Context) ([]domain.Booking, error) {
	bookings, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

// UpdateBooking updates a booking
func (s *service) UpdateBooking(ctx context.Context, booking domain.Booking) error {
	err := s.repo.Update(ctx, booking)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBooking deletes a booking
func (s *service) DeleteBooking(ctx context.Context, bookingID domain.BookingUUID) error {
	err := s.repo.Delete(ctx, bookingID)
	if err != nil {
		return err
	}
	return nil
}
