package app

import (
	"context"
	"gholi-fly-hotel/config"
	hotelPort "gholi-fly-hotel/internal/hotel/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	HotelService(ctx context.Context) hotelPort.Service
}
