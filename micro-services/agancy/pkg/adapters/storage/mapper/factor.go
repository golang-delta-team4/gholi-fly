package mapper

import (
	"gholi-fly-agancy/internal/factor/domain"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"gholi-fly-agancy/pkg/fp"

	"github.com/google/uuid"
)

// FactorDomain2Storage converts a Factor from the domain layer to the storage layer.
func FactorDomain2Storage(factorDomain domain.Factor) *types.Factor {
	return &types.Factor{
		ID:                uuid.UUID(factorDomain.ID),
		HotelFactorID:     factorDomain.HotelFactorID,
		TransportFactorID: factorDomain.TransportFactorID,
		ReservationID:     factorDomain.ReservationID,
		AgencyPrice:       factorDomain.AgencyPrice,
		Profit:            factorDomain.Profit,
		CreatedAt:         factorDomain.CreatedAt,
		UpdatedAt:         factorDomain.UpdatedAt,
	}
}

func factorDomain2Storage(factorDomain domain.Factor) types.Factor {
	return types.Factor{
		ID:                uuid.UUID(factorDomain.ID),
		HotelFactorID:     factorDomain.HotelFactorID,
		TransportFactorID: factorDomain.TransportFactorID,
		ReservationID:     factorDomain.ReservationID,
		AgencyPrice:       factorDomain.AgencyPrice,
		Profit:            factorDomain.Profit,
		CreatedAt:         factorDomain.CreatedAt,
		UpdatedAt:         factorDomain.UpdatedAt,
	}
}

func BatchFactorDomain2Storage(domains []domain.Factor) []types.Factor {
	return fp.Map(domains, factorDomain2Storage)
}

// FactorStorage2Domain converts a Factor from the storage layer to the domain layer.
func FactorStorage2Domain(factor types.Factor) *domain.Factor {
	return &domain.Factor{
		ID:                domain.FactorID(factor.ID),
		HotelFactorID:     factor.HotelFactorID,
		TransportFactorID: factor.TransportFactorID,
		ReservationID:     factor.ReservationID,
		AgencyPrice:       factor.AgencyPrice,
		Profit:            factor.Profit,
		CreatedAt:         factor.CreatedAt,
		UpdatedAt:         factor.UpdatedAt,
	}
}

func factorStorage2Domain(factor types.Factor) domain.Factor {
	return domain.Factor{
		ID:                domain.FactorID(factor.ID),
		HotelFactorID:     factor.HotelFactorID,
		TransportFactorID: factor.TransportFactorID,
		ReservationID:     factor.ReservationID,
		AgencyPrice:       factor.AgencyPrice,
		Profit:            factor.Profit,
		CreatedAt:         factor.CreatedAt,
		UpdatedAt:         factor.UpdatedAt,
	}
}

func BatchFactorStorage2Domain(factors []types.Factor) []domain.Factor {
	return fp.Map(factors, factorStorage2Domain)
}
