package booking

import (
	"context"
	"errors"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	"gholi-fly-hotel/internal/booking/port"
	roomDomain "gholi-fly-hotel/internal/room/domain"

	"github.com/google/uuid"
)

var (
	ErrBookingCreation           = errors.New("error on creating booking")
	ErrBookingCreationValidation = errors.New("error on creating booking: validation failed")
	ErrBookingCreationDuplicate  = errors.New("booking already exists")
	ErrBookingNotFound           = errors.New("booking not found")
	ErrInvalidSourceService      = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// CreateBookingByRoomID creates a new booking by room ID
func (s *service) CreateBookingByRoomID(ctx context.Context, booking bookingDomain.Booking, roomID roomDomain.RoomUUID) (bookingDomain.BookingUUID, error) {
	bookingID, err := s.repo.CreateByRoomID(ctx, booking, roomID)
	if err != nil {
		return bookingDomain.BookingUUID{}, ErrBookingCreation
	}
	return bookingID, nil
}

// GetAllBookingsByRoomID returns all bookings by room ID
func (s *service) GetAllBookingsByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]bookingDomain.Booking, error) {
	return s.repo.GetByRoomID(ctx, roomID)
}

// GetAllBookingsByUserID returns all bookings by user ID
func (s *service) GetAllBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]bookingDomain.Booking, error) {
	return s.repo.GetByUserID(ctx, userID)
}

// GetBookingByID returns a booking by its ID
func (s *service) GetBookingByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error) {
	booking, err := s.repo.GetByID(ctx, bookingID)
	if err != nil {
		return nil, ErrBookingNotFound
	}
	return booking, nil
}

// UpdateBooking updates a booking
func (s *service) UpdateBooking(ctx context.Context, booking bookingDomain.Booking) error {
	return s.repo.Update(ctx, booking)
}

// DeleteBooking deletes a booking
func (s *service) DeleteBooking(ctx context.Context, bookingID bookingDomain.BookingUUID) error {
	return s.repo.Delete(ctx, bookingID)
}
