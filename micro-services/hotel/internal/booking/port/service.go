package port

import (
	"context"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateBookingByHotelID(ctx context.Context, booking bookingDomain.Booking, hotelID hotelDomain.HotelUUID, isAgency bool) (bookingDomain.BookingUUID, roomDomain.RoomPrice, error)
	CreateBookingFactor(ctx context.Context, userId uuid.UUID, hotelID hotelDomain.HotelUUID, totalPrice uint, bookingId bookingDomain.BookingUUID) (string, error)
	GetAllBookingsByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]bookingDomain.Booking, error)
	GetAllBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]bookingDomain.Booking, error)
	GetBookingByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error)
	UpdateBooking(ctx context.Context, booking bookingDomain.Booking) error
	DeleteBooking(ctx context.Context, bookingID bookingDomain.BookingUUID) error
}
