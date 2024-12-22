package port

import (
	"context"
	roomDomain "gholi-fly-hotel/internal/room/domain"
)

type Repo interface {
	Create(ctx context.Context, room roomDomain.Room) (roomDomain.RoomUUID, error)
	GetByID(ctx context.Context, roomID roomDomain.RoomUUID) (*roomDomain.Room, error)
	GetAll(ctx context.Context) ([]roomDomain.Room, error)
	Update(ctx context.Context, room roomDomain.Room) error
	Delete(ctx context.Context, roomID roomDomain.RoomUUID) error
}
