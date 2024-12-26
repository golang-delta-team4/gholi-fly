package mapper

import (
	"gholi-fly-agancy/internal/agency/domain"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"gholi-fly-agancy/pkg/fp"

	"github.com/google/uuid"
)

// AgencyDomain2Storage converts an Agency from the domain layer to the storage layer.
func AgencyDomain2Storage(agencyDomain domain.Agency) *types.Agency {
	return &types.Agency{
		ID:               uuid.UUID(agencyDomain.ID),
		Name:             agencyDomain.Name,
		OwnerID:          agencyDomain.OwnerID,
		WalletID:         agencyDomain.WalletID,
		ProfitPercentage: agencyDomain.ProfitPercentage,
		CreatedAt:        agencyDomain.CreatedAt,
		UpdatedAt:        agencyDomain.UpdatedAt,
	}
}

func agencyDomain2Storage(agencyDomain domain.Agency) types.Agency {
	return types.Agency{
		ID:               uuid.UUID(agencyDomain.ID),
		Name:             agencyDomain.Name,
		OwnerID:          agencyDomain.OwnerID,
		WalletID:         agencyDomain.WalletID,
		ProfitPercentage: agencyDomain.ProfitPercentage,
		CreatedAt:        agencyDomain.CreatedAt,
		UpdatedAt:        agencyDomain.UpdatedAt,
	}
}

func BatchAgencyDomain2Storage(domains []domain.Agency) []types.Agency {
	return fp.Map(domains, agencyDomain2Storage)
}

// AgencyStorage2Domain converts an Agency from the storage layer to the domain layer.
func AgencyStorage2Domain(agency types.Agency) *domain.Agency {
	return &domain.Agency{
		ID:               domain.AgencyID(agency.ID),
		Name:             agency.Name,
		OwnerID:          agency.OwnerID,
		WalletID:         agency.WalletID,
		ProfitPercentage: agency.ProfitPercentage,
		CreatedAt:        agency.CreatedAt,
		UpdatedAt:        agency.UpdatedAt,
	}
}

func agencyStorage2Domain(agency types.Agency) domain.Agency {
	return domain.Agency{
		ID:               domain.AgencyID(agency.ID),
		Name:             agency.Name,
		OwnerID:          agency.OwnerID,
		WalletID:         agency.WalletID,
		ProfitPercentage: agency.ProfitPercentage,
		CreatedAt:        agency.CreatedAt,
		UpdatedAt:        agency.UpdatedAt,
	}
}

func BatchAgencyStorage2Domain(agencies []types.Agency) []domain.Agency {
	return fp.Map(agencies, agencyStorage2Domain)
}
