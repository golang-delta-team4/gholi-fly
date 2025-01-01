package mapper

import (
	"gholi-fly-agancy/internal/staff/domain"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"gholi-fly-agancy/pkg/fp"

	"github.com/google/uuid"
)

// StaffDomain2Storage converts a Staff from the domain layer to the storage layer.
func StaffDomain2Storage(staffDomain domain.Staff) *types.Staff {
	return &types.Staff{
		ID:        uuid.UUID(staffDomain.ID),
		UserID:    staffDomain.UserID,
		AgencyID:  staffDomain.AgencyID,
		WalletID:  staffDomain.WalletID,
		Stock:     staffDomain.Stock,
		Role:      staffDomain.Role,
		CreatedAt: staffDomain.CreatedAt,
		UpdatedAt: staffDomain.UpdatedAt,
	}
}

func staffDomain2Storage(staffDomain domain.Staff) types.Staff {
	return types.Staff{
		ID:        uuid.UUID(staffDomain.ID),
		UserID:    staffDomain.UserID,
		AgencyID:  staffDomain.AgencyID,
		WalletID:  staffDomain.WalletID,
		Stock:     staffDomain.Stock,
		Role:      staffDomain.Role,
		CreatedAt: staffDomain.CreatedAt,
		UpdatedAt: staffDomain.UpdatedAt,
	}
}

func BatchStaffDomain2Storage(domains []domain.Staff) []types.Staff {
	return fp.Map(domains, staffDomain2Storage)
}

// StaffStorage2Domain converts a Staff from the storage layer to the domain layer.
func StaffStorage2Domain(staff types.Staff) *domain.Staff {
	return &domain.Staff{
		ID:        domain.StaffID(staff.ID),
		UserID:    staff.UserID,
		AgencyID:  staff.AgencyID,
		WalletID:  staff.WalletID,
		Stock:     staff.Stock,
		Role:      staff.Role,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
	}
}

func staffStorage2Domain(staff types.Staff) domain.Staff {
	return domain.Staff{
		ID:        domain.StaffID(staff.ID),
		UserID:    staff.UserID,
		AgencyID:  staff.AgencyID,
		WalletID:  staff.WalletID,
		Stock:     staff.Stock,
		Role:      staff.Role,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
	}
}

func BatchStaffStorage2Domain(staffs []types.Staff) []domain.Staff {
	return fp.Map(staffs, staffStorage2Domain)
}
