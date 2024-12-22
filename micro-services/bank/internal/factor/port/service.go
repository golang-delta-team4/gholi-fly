package port

import (
	"context"
	"gholi-fly-bank/internal/factor/domain"
)

type Service interface {
	// Create a new factor in the system.
	CreateFactor(ctx context.Context, factor domain.Factor) (domain.FactorUUID, error)

	// Retrieve a factor by its ID.
	GetFactorByID(ctx context.Context, factorID domain.FactorUUID) (*domain.Factor, error)

	// Retrieve factors based on filters.
	GetFactors(ctx context.Context, filters domain.FactorFilters) ([]domain.Factor, error)

	// Update the status of a factor.
	UpdateFactorStatus(ctx context.Context, factorID domain.FactorUUID, status domain.FactorStatus) error
}
