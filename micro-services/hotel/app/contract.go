package app

import (
	"gholi-fly-hotel/config"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
}
