package mapper

import (
	"gholi-fly-hotel/internal/room/domain"
	"gholi-fly-hotel/pkg/adapters/storage/types"
	"gholi-fly-hotel/pkg/fp"

	"gorm.io/gorm"
)

func RoomDomain2Storage(roomDomain domain.Room) *types.Room {
	return &types.Room{
		Model: gorm.Model{
			CreatedAt: roomDomain.CreatedAt,
			UpdatedAt: roomDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(roomDomain.DeletedAt)),
		},
		UUID:        roomDomain.UUID,
		HotelID:     roomDomain.HotelID,
		RoomNumber:  roomDomain.RoomNumber,
		Floor:       roomDomain.Floor,
		BasePrice:   roomDomain.BasePrice,
		AgencyPrice: roomDomain.AgencyPrice,
	}
}

func roomDomain2Storage(roomDomain domain.Room) types.Room {
	return types.Room{
		Model: gorm.Model{
			CreatedAt: roomDomain.CreatedAt,
			UpdatedAt: roomDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(roomDomain.DeletedAt)),
		},
		UUID:        roomDomain.UUID,
		HotelID:     roomDomain.HotelID,
		RoomNumber:  roomDomain.RoomNumber,
		Floor:       roomDomain.Floor,
		BasePrice:   roomDomain.BasePrice,
		AgencyPrice: roomDomain.AgencyPrice,
	}
}

func BatchRoomDomain2Storage(domains []domain.Room) []types.Room {
	return fp.Map(domains, roomDomain2Storage)
}

func RoomStorage2Domain(room types.Room) *domain.Room {
	return &domain.Room{
		ID:          domain.RoomID(room.ID),
		UUID:        domain.RoomUUID(room.UUID),
		HotelID:     room.HotelID,
		RoomNumber:  room.RoomNumber,
		Floor:       room.Floor,
		BasePrice:   room.BasePrice,
		AgencyPrice: room.AgencyPrice,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
		DeletedAt:   room.DeletedAt.Time,
	}
}

func roomStorage2Domain(room types.Room) domain.Room {
	return domain.Room{
		ID:          domain.RoomID(room.ID),
		UUID:        domain.RoomUUID(room.UUID),
		HotelID:     room.HotelID,
		RoomNumber:  room.RoomNumber,
		Floor:       room.Floor,
		BasePrice:   room.BasePrice,
		AgencyPrice: room.AgencyPrice,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
		DeletedAt:   room.DeletedAt.Time,
	}
}

func BatchRoomStorage2Domain(rooms []types.Room) []domain.Room {
	return fp.Map(rooms, roomStorage2Domain)
}
