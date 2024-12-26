package port

import (
	"context"

	"gholi-fly-agancy/internal/staff/domain"

	"github.com/google/uuid"
)

type StaffService interface {
	CreateStaff(ctx context.Context, staff domain.Staff) (domain.StaffID, error)
	GetStaffByID(ctx context.Context, id domain.StaffID) (*domain.Staff, error)
	UpdateStaff(ctx context.Context, staff domain.Staff) error
	DeleteStaff(ctx context.Context, id domain.StaffID) error
	ListStaffByAgency(ctx context.Context, agencyID uuid.UUID) ([]domain.Staff, error)
}
