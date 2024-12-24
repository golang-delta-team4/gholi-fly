package port

import (
	"context"
	"gholi-fly-bank/internal/factor/domain"
)

type Repo interface {
	// Create a new factor in the system.
	Create(ctx context.Context, factor domain.Factor) (domain.FactorUUID, error)

	// Retrieve a factor by its ID.
	GetByID(ctx context.Context, factorID domain.FactorUUID) (*domain.Factor, error)

	// Retrieve factors based on filters with pagination.
	Get(ctx context.Context, filters domain.FactorFilters) ([]domain.Factor, int, error)

	// Update the status of a factor.
	UpdateStatus(ctx context.Context, factorID domain.FactorUUID, status domain.FactorStatus) error
}
