package booking

import (
	"context"
	"errors"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	"gholi-fly-hotel/internal/booking/port"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrBookingCreation           = errors.New("error on creating booking")
	ErrBookingCreationValidation = errors.New("error on creating booking: validation failed")
	ErrBookingCreationDuplicate  = errors.New("booking already exists in these days")
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
func (s *service) CreateBookingByHotelID(ctx context.Context, booking bookingDomain.Booking, hotelID hotelDomain.HotelUUID) (bookingDomain.BookingUUID, roomDomain.RoomPrice, error) {
	if err := booking.Validate(); err != nil {
		return uuid.Nil, 0, ErrBookingCreationValidation
	}
	bookingID, price, err := s.repo.CreateByHotelID(ctx, booking, hotelID)
	if err != nil {
		if strings.Contains(err.Error(), ErrBookingCreationDuplicate.Error()) {
			return bookingDomain.BookingUUID{}, 0, ErrBookingCreationDuplicate
		}
		return bookingDomain.BookingUUID{}, 0, ErrBookingCreation
	}
	return bookingID, price, nil
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
