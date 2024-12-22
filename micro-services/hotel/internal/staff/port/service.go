package port

import (
	"context"
	staffDomain "gholi-fly-hotel/internal/staff/domain"
)

type Service interface {
	CreateStaff(ctx context.Context, staff staffDomain.Staff) (staffDomain.StaffUUID, error)
	GetStaffByID(ctx context.Context, staffID staffDomain.StaffUUID) (*staffDomain.Staff, error)
	GetStaffs(ctx context.Context) ([]staffDomain.Staff, error)
	UpdateStaff(ctx context.Context, staff staffDomain.Staff) error
	DeleteStaff(ctx context.Context, staffID staffDomain.StaffUUID) error
}
