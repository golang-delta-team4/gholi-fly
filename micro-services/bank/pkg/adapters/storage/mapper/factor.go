package mapper

import (
	"gholi-fly-bank/internal/factor/domain"
	"gholi-fly-bank/pkg/adapters/storage/types"
	"gholi-fly-bank/pkg/fp"

	"github.com/google/uuid"
)

// FactorDomain2Storage converts a Factor from the domain layer to the storage layer.
func FactorDomain2Storage(factorDomain domain.Factor) *types.Factor {
	return &types.Factor{
		ID:            uuid.UUID(factorDomain.ID), // Cast domain ID to storage ID.
		SourceService: factorDomain.SourceService,
		ExternalID:    factorDomain.ExternalID,
		BookingID:     factorDomain.BookingID,
		Amount:        factorDomain.Amount,
		CustomerID:    factorDomain.CustomerID,
		Status:        uint8(factorDomain.Status), // Cast domain.FactorStatus to uint8.
		Details:       factorDomain.Details,
		CreatedAt:     factorDomain.CreatedAt,
		UpdatedAt:     factorDomain.UpdatedAt,
	}
}

func factorDomain2Storage(factorDomain domain.Factor) types.Factor {
	return types.Factor{
		ID:            uuid.UUID(factorDomain.ID), // Cast domain ID to storage ID.
		SourceService: factorDomain.SourceService,
		ExternalID:    factorDomain.ExternalID,
		BookingID:     factorDomain.BookingID,
		Amount:        factorDomain.Amount,
		CustomerID:    factorDomain.CustomerID,
		Status:        uint8(factorDomain.Status), // Cast domain.FactorStatus to uint8.
		Details:       factorDomain.Details,
		CreatedAt:     factorDomain.CreatedAt,
		UpdatedAt:     factorDomain.UpdatedAt,
	}
}

func BatchFactorDomain2Storage(domains []domain.Factor) []types.Factor {
	return fp.Map(domains, factorDomain2Storage)
}

// FactorStorage2Domain converts a Factor from the storage layer to the domain layer.
func FactorStorage2Domain(factor types.Factor) *domain.Factor {
	return &domain.Factor{
		ID:            domain.FactorUUID(factor.ID), // Cast storage ID to domain ID.
		SourceService: factor.SourceService,
		ExternalID:    factor.ExternalID,
		BookingID:     factor.BookingID,
		Amount:        factor.Amount,
		CustomerID:    factor.CustomerID,
		Status:        domain.FactorStatus(factor.Status), // Cast uint8 to domain.FactorStatus.
		Details:       factor.Details,
		CreatedAt:     factor.CreatedAt,
		UpdatedAt:     factor.UpdatedAt,
	}
}

func factorStorage2Domain(factor types.Factor) domain.Factor {
	return domain.Factor{
		ID:            domain.FactorUUID(factor.ID), // Cast storage ID to domain ID.
		SourceService: factor.SourceService,
		ExternalID:    factor.ExternalID,
		BookingID:     factor.BookingID,
		Amount:        factor.Amount,
		CustomerID:    factor.CustomerID,
		Status:        domain.FactorStatus(factor.Status), // Cast uint8 to domain.FactorStatus.
		Details:       factor.Details,
		CreatedAt:     factor.CreatedAt,
		UpdatedAt:     factor.UpdatedAt,
	}
}

func BatchFactorStorage2Domain(factors []types.Factor) []domain.Factor {
	return fp.Map(factors, factorStorage2Domain)
}
