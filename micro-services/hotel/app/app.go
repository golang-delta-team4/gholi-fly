package app

import (
	"context"
	"fmt"
	"gholi-fly-hotel/config"
	"gholi-fly-hotel/internal/hotel"
	hotelPort "gholi-fly-hotel/internal/hotel/port"
	"gholi-fly-hotel/internal/room"
	roomPort "gholi-fly-hotel/internal/room/port"
	"gholi-fly-hotel/internal/staff"
	staffPort "gholi-fly-hotel/internal/staff/port"
	"gholi-fly-hotel/pkg/adapters/storage"
	"gholi-fly-hotel/pkg/postgres"

	"gorm.io/gorm"

	appCtx "gholi-fly-hotel/pkg/context"
)

type app struct {
	db           *gorm.DB
	cfg          config.Config
	hotelService hotelPort.Service
	roomService  roomPort.Service
	staffService staffPort.Service
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

func (a *app) RoomService(ctx context.Context) roomPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.roomService == nil {
			a.roomService = a.roomServiceWithDB(a.db)
		}
		return a.roomService
	}

	return a.roomServiceWithDB(db)
}

func (a *app) roomServiceWithDB(db *gorm.DB) roomPort.Service {
	return room.NewService(storage.NewRoomRepo(db))
}

func (a *app) StaffService(ctx context.Context) staffPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.staffService == nil {
			a.staffService = a.staffServiceWithDB(a.db)
		}
		return a.staffService
	}

	return a.staffServiceWithDB(db)
}

func (a *app) staffServiceWithDB(db *gorm.DB) staffPort.Service {
	return staff.NewService(storage.NewStaffRepo(db))
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
