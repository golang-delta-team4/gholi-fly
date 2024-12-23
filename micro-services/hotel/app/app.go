package app

import (
	"context"
	"fmt"
	"gholi-fly-hotel/config"
	"gholi-fly-hotel/internal/hotel"
	hotelPort "gholi-fly-hotel/internal/hotel/port"
	"gholi-fly-hotel/pkg/adapters/storage"
	"gholi-fly-hotel/pkg/postgres"

	"gorm.io/gorm"

	appCtx "gholi-fly-hotel/pkg/context"
)

type app struct {
	db           *gorm.DB
	cfg          config.Config
	hotelService hotelPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) HotelService(ctx context.Context) hotelPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.hotelService == nil {
			a.hotelService = a.hotelServiceWithDB(a.db)
		}
		return a.hotelService
	}

	return a.hotelServiceWithDB(db)
}

func (a *app) hotelServiceWithDB(db *gorm.DB) hotelPort.Service {
	return hotel.NewService(storage.NewHotelRepo(db))
}

func (a *app) Config() config.Config {
	return a.cfg
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

	if err := postgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	a.db = db
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
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
