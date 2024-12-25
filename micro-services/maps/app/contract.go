package app

import (
	"gholi-fly-maps/config"
	port_p "gholi-fly-maps/internal/paths/port"
	"gholi-fly-maps/internal/terminals/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	TerminalService() port.TerminalService
	PathService() port_p.PathService
}
