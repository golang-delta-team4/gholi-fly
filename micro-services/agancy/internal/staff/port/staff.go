package port

import (
	"context"

	"gholi-fly-agancy/internal/staff/domain"

	"github.com/google/uuid"
)

type StaffRepo interface {
	Create(ctx context.Context, staff domain.Staff) (domain.StaffID, error)
	GetByID(ctx context.Context, id domain.StaffID) (*domain.Staff, error)
	Update(ctx context.Context, staff domain.Staff) error
	Delete(ctx context.Context, id domain.StaffID) error
	ListByAgencyID(ctx context.Context, agencyID uuid.UUID) ([]domain.Staff, error)
}
