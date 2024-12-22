package port

import (
	"context"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
)

type Repo interface {
	Create(ctx context.Context, booking bookingDomain.Booking) (bookingDomain.BookingUUID, error)
	GetByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error)
	GetAll(ctx context.Context) ([]bookingDomain.Booking, error)
	Update(ctx context.Context, booking bookingDomain.Booking) error
	Delete(ctx context.Context, bookingID bookingDomain.BookingUUID) error
}
