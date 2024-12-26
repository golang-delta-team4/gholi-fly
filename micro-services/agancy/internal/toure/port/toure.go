package port

import (
	"context"

	"gholi-fly-agancy/internal/toure/domain"

	"github.com/google/uuid"
)

type ToureRepo interface {
	Create(ctx context.Context, toure domain.Toure) (domain.ToureID, error)
	GetByID(ctx context.Context, id domain.ToureID) (*domain.Toure, error)
	Update(ctx context.Context, toure domain.Toure) error
	Delete(ctx context.Context, id domain.ToureID) error
	ListByAgencyID(ctx context.Context, agencyID uuid.UUID) ([]domain.Toure, error)
}
