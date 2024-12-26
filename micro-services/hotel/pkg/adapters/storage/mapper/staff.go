package mapper

import (
	"gholi-fly-hotel/internal/staff/domain"
	"gholi-fly-hotel/pkg/adapters/storage/types"
	"gholi-fly-hotel/pkg/fp"

	"gorm.io/gorm"
)

func StaffDomain2Storage(staffDomain domain.Staff) *types.Staff {
	return &types.Staff{
		Model: gorm.Model{
			CreatedAt: staffDomain.CreatedAt,
			UpdatedAt: staffDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(staffDomain.DeletedAt)),
		},
		UUID:      staffDomain.UUID,
		HotelID:   staffDomain.HotelID,
		Name:      staffDomain.Name,
		StaffType: staffDomain.StaffType,
	}
}

func staffDomain2Storage(staffDomain domain.Staff) types.Staff {
	return types.Staff{
		Model: gorm.Model{
			CreatedAt: staffDomain.CreatedAt,
			UpdatedAt: staffDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(staffDomain.DeletedAt)),
		},
		UUID:      staffDomain.UUID,
		HotelID:   staffDomain.HotelID,
		Name:      staffDomain.Name,
		StaffType: staffDomain.StaffType,
	}
}

func BatchStaffDomain2Storage(domains []domain.Staff) []types.Staff {
	return fp.Map(domains, staffDomain2Storage)
}

func StaffStorage2Domain(staff types.Staff) *domain.Staff {
	return &domain.Staff{
		ID:        domain.StaffID(staff.ID),
		UUID:      domain.StaffUUID(staff.UUID),
		HotelID:   staff.HotelID,
		Name:      staff.Name,
		StaffType: staff.StaffType,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
		DeletedAt: staff.DeletedAt.Time,
	}
}

func staffStorage2Domain(staff types.Staff) domain.Staff {
	return domain.Staff{
		ID:        domain.StaffID(staff.ID),
		UUID:      domain.StaffUUID(staff.UUID),
		HotelID:   staff.HotelID,
		Name:      staff.Name,
		StaffType: staff.StaffType,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
		DeletedAt: staff.DeletedAt.Time,
	}
}

func BatchStaffStorage2Domain(staffs []types.Staff) []domain.Staff {
	return fp.Map(staffs, staffStorage2Domain)
}
