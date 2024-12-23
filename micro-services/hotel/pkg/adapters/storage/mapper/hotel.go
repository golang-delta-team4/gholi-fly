package mapper

import (
	"gholi-fly-hotel/internal/hotel/domain"
	"gholi-fly-hotel/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func HotelDomain2Storage(hotelDomain domain.Hotel) *types.Hotel {
	return &types.Hotel{
		Model: gorm.Model{
			ID:        uint(hotelDomain.ID),
			CreatedAt: hotelDomain.CreatedAt,
			UpdatedAt: hotelDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(hotelDomain.DeletedAt)),
		},
		UUID:       hotelDomain.UUID,
		OwnerEmail: hotelDomain.OwnerEmail,
		Name:       hotelDomain.Name,
		City:       hotelDomain.City,
	}
}

func HotelStorage2Domain(hotel types.Hotel) *domain.Hotel {
	// uid, err := domain.HotelUUIDFromString(hotel.UUID)
	return &domain.Hotel{
		ID: domain.HotelID(hotel.ID),
		// UUID:       uid,
		UUID:       domain.HotelUUID(hotel.UUID),
		OwnerEmail: hotel.OwnerEmail,
		Name:       hotel.Name,
		City:       hotel.City,
		CreatedAt:  hotel.CreatedAt,
		UpdatedAt:  hotel.UpdatedAt,
		DeletedAt:  hotel.DeletedAt.Time,
	}
}
