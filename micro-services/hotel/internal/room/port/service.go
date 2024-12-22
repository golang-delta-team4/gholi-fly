package port

import (
	"context"
	roomDomain "gholi-fly-hotel/internal/room/domain"
)

type Service interface {
	CreateRoom(ctx context.Context, room roomDomain.Room) (roomDomain.RoomUUID, error)
	GetRoomByID(ctx context.Context, roomID roomDomain.RoomUUID) (*roomDomain.Room, error)
	GetRooms(ctx context.Context) ([]roomDomain.Room, error)
	UpdateRoom(ctx context.Context, room roomDomain.Room) error
	DeleteRoom(ctx context.Context, roomID roomDomain.RoomUUID) error
}
