package mapper

import (
	"gholi-fly-hotel/internal/hotel/domain"
	"gholi-fly-hotel/pkg/adapters/storage/types"
	"gholi-fly-hotel/pkg/fp"

	"gorm.io/gorm"
)

func HotelDomain2Storage(hotelDomain domain.Hotel) *types.Hotel {
	return &types.Hotel{
		Model: gorm.Model{
			CreatedAt: hotelDomain.CreatedAt,
			UpdatedAt: hotelDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(hotelDomain.DeletedAt)),
		},
		UUID:    hotelDomain.UUID,
		OwnerID: hotelDomain.OwnerID,
		Name:    hotelDomain.Name,
		City:    hotelDomain.City,
	}
}

func hotelDomain2Storage(hotelDomain domain.Hotel) types.Hotel {
	return types.Hotel{
		Model: gorm.Model{
			CreatedAt: hotelDomain.CreatedAt,
			UpdatedAt: hotelDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(hotelDomain.DeletedAt)),
		},
		UUID:    hotelDomain.UUID,
		OwnerID: hotelDomain.OwnerID,
		Name:    hotelDomain.Name,
		City:    hotelDomain.City,
	}
}

func BatchHotelDomain2Storage(domains []domain.Hotel) []types.Hotel {
	return fp.Map(domains, hotelDomain2Storage)
}

func HotelStorage2Domain(hotel types.Hotel) *domain.Hotel {
	return &domain.Hotel{
		UUID:      domain.HotelUUID(hotel.UUID),
		OwnerID:   hotel.OwnerID,
		Name:      hotel.Name,
		City:      hotel.City,
		CreatedAt: hotel.CreatedAt,
		UpdatedAt: hotel.UpdatedAt,
		DeletedAt: hotel.DeletedAt.Time,
	}
}

func hotelStorage2Domain(hotel types.Hotel) domain.Hotel {
	return domain.Hotel{
		UUID:      domain.HotelUUID(hotel.UUID),
		OwnerID:   hotel.OwnerID,
		Name:      hotel.Name,
		City:      hotel.City,
		CreatedAt: hotel.CreatedAt,
		UpdatedAt: hotel.UpdatedAt,
		DeletedAt: hotel.DeletedAt.Time,
	}
}

func BatchHotelStorage2Domain(hotels []types.Hotel) []domain.Hotel {
	return fp.Map(hotels, hotelStorage2Domain)
}
