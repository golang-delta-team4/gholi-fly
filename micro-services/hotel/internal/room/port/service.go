package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"
)

type Service interface {
	CreateRoomByHotelID(ctx context.Context, room roomDomain.Room, hotelID hotelDomain.HotelUUID) (roomDomain.RoomUUID, error)
	GetAllRoomsByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]roomDomain.Room, error)
	GetRoomByID(ctx context.Context, roomID roomDomain.RoomUUID) (*roomDomain.Room, error)
	UpdateRoom(ctx context.Context, room roomDomain.Room) error
	DeleteRoom(ctx context.Context, roomID roomDomain.RoomUUID) error
}
