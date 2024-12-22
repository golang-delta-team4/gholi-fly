package port

import (
	"context"
	staffDomain "gholi-fly-hotel/internal/staff/domain"
)

type Repo interface {
	Create(ctx context.Context, staff staffDomain.Staff) (staffDomain.StaffUUID, error)
	GetByID(ctx context.Context, staffID staffDomain.StaffUUID) (*staffDomain.Staff, error)
	GetAll(ctx context.Context) ([]staffDomain.Staff, error)
	Update(ctx context.Context, staff staffDomain.Staff) error
	Delete(ctx context.Context, staffID staffDomain.StaffUUID) error
}
