package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	staffDomain "gholi-fly-hotel/internal/staff/domain"
)

type Service interface {
	CreateStaffByHotelID(ctx context.Context, staff staffDomain.Staff, hotelID hotelDomain.HotelUUID) (staffDomain.StaffUUID, error)
	GetAllStaffsByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]staffDomain.Staff, error)
	GetStaffByID(ctx context.Context, staffID staffDomain.StaffUUID) (*staffDomain.Staff, error)
	UpdateStaff(ctx context.Context, staff staffDomain.Staff) error
	DeleteStaff(ctx context.Context, staffID staffDomain.StaffUUID) error
}
