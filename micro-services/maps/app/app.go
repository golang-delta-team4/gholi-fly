package app

import (
	"fmt"
	"gholi-fly-maps/config"
	"gholi-fly-maps/internal/paths"
	path_domain "gholi-fly-maps/internal/paths/domain"
	port_p "gholi-fly-maps/internal/paths/port"
	"gholi-fly-maps/internal/terminals"
	terminal_domain "gholi-fly-maps/internal/terminals/domain"
	port_t "gholi-fly-maps/internal/terminals/port"
	"gholi-fly-maps/pkg/adapters/storage"
	"gholi-fly-maps/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db              *gorm.DB
	cfg             config.Config
	terminalService port_t.TerminalService
	pathService     port_p.PathService
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) TerminalService() port_t.TerminalService {
	return a.terminalService
}
func (a *app) PathService() port_p.PathService {
	return a.pathService
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})
	if err != nil {
		return err
	}
	// Enable the uuid-ossp extension
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return fmt.Errorf("failed to create uuid-ossp extension: %v", err)
	}

	// Run migrations for Terminal and Path models
	err = postgres.AutoMigrate(db, &terminal_domain.Terminal{}, &path_domain.Path{})
	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setServices() error {
	// Initialize repositories
	terminalRepo := storage.NewTerminalRepo(a.db)
	pathRepo := storage.NewPathRepo(a.db)

	// Initialize services
	a.terminalService = terminals.NewTerminalService(terminalRepo)
	a.pathService = paths.NewPathService(pathRepo, terminalRepo)

	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	// Set up the database
	if err := a.setDB(); err != nil {
		return nil, err
	}

	// Set up the services
	if err := a.setServices(); err != nil {
		return nil, err
	}

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
