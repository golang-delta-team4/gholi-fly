package app

import (
	"context"
	"fmt"

	"gholi-fly-agancy/config"
	"gholi-fly-agancy/internal/agency"
	agencyPort "gholi-fly-agancy/internal/agency/port"
	"gholi-fly-agancy/internal/factor"
	factorPort "gholi-fly-agancy/internal/factor/port"
	"gholi-fly-agancy/internal/reservation"
	reservationPort "gholi-fly-agancy/internal/reservation/port"
	"gholi-fly-agancy/internal/staff"
	staffPort "gholi-fly-agancy/internal/staff/port"
	"gholi-fly-agancy/internal/tour"
	tourPort "gholi-fly-agancy/internal/tour/port"
	"gholi-fly-agancy/pkg/adapters/storage"
	"gholi-fly-agancy/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db                 *gorm.DB
	cfg                config.Config
	agencyService      agencyPort.AgencyService
	staffService       staffPort.StaffService
	factorService      factorPort.FactorService
	tourService        tourPort.TourService
	reservationService reservationPort.ReservationService
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) AgencyService(ctx context.Context) agencyPort.AgencyService {
	if a.agencyService == nil {
		a.agencyService = agency.NewService(storage.NewAgencyRepo(a.db))
	}
	return a.agencyService
}

func (a *app) StaffService(ctx context.Context) staffPort.StaffService {
	if a.staffService == nil {
		a.staffService = staff.NewService(storage.NewStaffRepo(a.db))
	}
	return a.staffService
}

func (a *app) FactorService(ctx context.Context) factorPort.FactorService {
	if a.factorService == nil {
		a.factorService = factor.NewService(storage.NewFactorRepo(a.db))
	}
	return a.factorService
}

func (a *app) TourService(ctx context.Context) tourPort.TourService {
	if a.tourService == nil {
		a.tourService = tour.NewService(storage.NewTourRepo(a.db))
	}
	return a.tourService
}

func (a *app) ReservationService(ctx context.Context) reservationPort.ReservationService {
	if a.reservationService == nil {
		a.reservationService = reservation.NewService(storage.NewReservationRepo(a.db))
	}
	return a.reservationService
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
	})
	if err != nil {
		return err
	}
	a.db = db
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{cfg: cfg}
	if err := a.setDB(); err != nil {
		return nil, err
	}
	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app: %v", err))
	}
	return app
}
