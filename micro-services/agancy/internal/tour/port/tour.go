package port

import (
	"context"
	"gholi-fly-agancy/internal/tour/domain"

	"github.com/google/uuid"
)

type TourRepo interface {
	Create(ctx context.Context, tour domain.Tour) (domain.TourID, error)
	GetByID(ctx context.Context, id domain.TourID) (*domain.Tour, error)
	Update(ctx context.Context, tour domain.Tour) error
	Delete(ctx context.Context, id domain.TourID) error
	ListByAgencyID(ctx context.Context, agencyID uuid.UUID) ([]domain.Tour, error)
}
