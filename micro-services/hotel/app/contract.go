package app

import (
	"context"
	"gholi-fly-hotel/config"
	hotelPort "gholi-fly-hotel/internal/hotel/port"
	roomPort "gholi-fly-hotel/internal/room/port"
	staffPort "gholi-fly-hotel/internal/staff/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	HotelService(ctx context.Context) hotelPort.Service
	RoomService(ctx context.Context) roomPort.Service
	StaffService(ctx context.Context) staffPort.Service
}
