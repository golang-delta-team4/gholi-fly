package port

import (
	"context"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
)

type Service interface {
	CreateBooking(ctx context.Context, booking bookingDomain.Booking) (bookingDomain.BookingUUID, error)
	GetBookingByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error)
	GetBookings(ctx context.Context) ([]bookingDomain.Booking, error)
	UpdateBooking(ctx context.Context, booking bookingDomain.Booking) error
	DeleteBooking(ctx context.Context, bookingID bookingDomain.BookingUUID) error
}
