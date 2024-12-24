package port

import (
	"context"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateBookingByRoomID(ctx context.Context, booking bookingDomain.Booking, roomID roomDomain.RoomUUID) (bookingDomain.BookingUUID, error)
	GetAllBookingsByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]bookingDomain.Booking, error)
	GetAllBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]bookingDomain.Booking, error)
	GetBookingByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error)
	UpdateBooking(ctx context.Context, booking bookingDomain.Booking) error
	DeleteBooking(ctx context.Context, bookingID bookingDomain.BookingUUID) error
}
