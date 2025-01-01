package port

import (
	"context"

	"gholi-fly-agancy/internal/factor/domain"

	"github.com/google/uuid"
)

type FactorRepo interface {
	Create(ctx context.Context, factor domain.Factor) (domain.FactorID, error)
	GetByID(ctx context.Context, id domain.FactorID) (*domain.Factor, error)
	Update(ctx context.Context, factor domain.Factor) error
	Delete(ctx context.Context, id domain.FactorID) error
	ListByReservationID(ctx context.Context, reservationID uuid.UUID) ([]domain.Factor, error)
}
