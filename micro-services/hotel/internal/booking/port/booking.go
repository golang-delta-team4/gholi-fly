package port

import (
	"context"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"

	"github.com/google/uuid"
)

type Repo interface {
	CreateByHotelID(ctx context.Context, booking bookingDomain.Booking, hotelID hotelDomain.HotelUUID, isAgency bool) (bookingDomain.BookingUUID, roomDomain.RoomPrice, error)
	GetByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]bookingDomain.Booking, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]bookingDomain.Booking, error)
	GetAllBookingsByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]bookingDomain.Booking, error)
	GetByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error)
	Update(ctx context.Context, booking bookingDomain.Booking) error
	AddBookingFactor(ctx context.Context, bookingID bookingDomain.BookingUUID, factorID string) error
	Delete(ctx context.Context, bookingID bookingDomain.BookingUUID) error
	ApproveUserBooking(ctx context.Context, factorID uuid.UUID, userUUID uuid.UUID) error
	CancelUserBooking(ctx context.Context, factorID uuid.UUID, userUUID uuid.UUID) error
	CancelBooking(ctx context.Context, factorID uuid.UUID) error
}
