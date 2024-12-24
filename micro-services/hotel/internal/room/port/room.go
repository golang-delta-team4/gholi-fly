package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"
)

type Repo interface {
	CreateByHotelID(ctx context.Context, room roomDomain.Room, hotelID hotelDomain.HotelUUID) (roomDomain.RoomUUID, error)
	GetByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]roomDomain.Room, error)
	GetByID(ctx context.Context, roomID roomDomain.RoomUUID) (*roomDomain.Room, error)
	Update(ctx context.Context, room roomDomain.Room) error
	Delete(ctx context.Context, roomID roomDomain.RoomUUID) error
}
