package port

import (
	"context"

	"gholi-fly-agancy/internal/factor/domain"

	"github.com/google/uuid"
)

type FactorService interface {
	CreateFactor(ctx context.Context, factor domain.Factor) (domain.FactorID, error)
	GetFactorByID(ctx context.Context, id domain.FactorID) (*domain.Factor, error)
	UpdateFactor(ctx context.Context, factor domain.Factor) error
	DeleteFactor(ctx context.Context, id domain.FactorID) error
	ListFactorsByReservation(ctx context.Context, reservationID uuid.UUID) ([]domain.Factor, error)
}
