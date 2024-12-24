package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	staffDomain "gholi-fly-hotel/internal/staff/domain"
)

type Repo interface {
	CreateByHotelID(ctx context.Context, staff staffDomain.Staff, hotelID hotelDomain.HotelUUID) (staffDomain.StaffUUID, error)
	GetByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]staffDomain.Staff, error)
	GetByID(ctx context.Context, staffID staffDomain.StaffUUID) (*staffDomain.Staff, error)
	Update(ctx context.Context, staff staffDomain.Staff) error
	Delete(ctx context.Context, staffID staffDomain.StaffUUID) error
}
