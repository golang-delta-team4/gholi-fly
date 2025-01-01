package app

import (
	"vehicle/config"
	"vehicle/internal/vehicle/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	VehicleService() port.VehicleService
}
