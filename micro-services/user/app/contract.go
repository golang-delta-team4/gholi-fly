package app

import (
	"user-service/config"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
}
