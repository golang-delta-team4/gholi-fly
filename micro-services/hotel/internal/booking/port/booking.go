package port

import (
	"context"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"

	"github.com/google/uuid"
)

type Repo interface {
	CreateByRoomID(ctx context.Context, booking bookingDomain.Booking, roomID roomDomain.RoomUUID) (bookingDomain.BookingUUID, error)
	GetByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]bookingDomain.Booking, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]bookingDomain.Booking, error)
	GetByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error)
	Update(ctx context.Context, booking bookingDomain.Booking) error
	Delete(ctx context.Context, bookingID bookingDomain.BookingUUID) error
}
