package app

import (
	"fmt"
	"vehicle/config"
	"vehicle/internal/vehicle"
	"vehicle/internal/vehicle/domain"
	"vehicle/internal/vehicle/port"
	"vehicle/pkg/adapters/storage"
	"vehicle/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db             *gorm.DB
	cfg            config.Config
	vehicleService port.VehicleService
	
	
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) VehicleService() port.VehicleService {
	return a.vehicleService
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

	// Run migrations
	err = postgres.AutoMigrate(db, &domain.Vehicle{}, &domain.TripRequest{})
	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setServices() error {
    vehicleRepo := storage.NewVehicleRepo(a.db)
    a.vehicleService = vehicle.NewVehicleService(vehicleRepo, a.cfg.TripService.URL) 
	return nil
}

func NewApp(cfg config.Config) (*app, error) {
	a := &app{cfg: cfg}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	if err := a.setServices(); err != nil {
		return nil, err
	}

	return a, nil
}
