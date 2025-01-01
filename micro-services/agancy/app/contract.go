package app

import (
	"context"
	config "gholi-fly-agancy/config"
	agencyPort "gholi-fly-agancy/internal/agency/port"
	factorPort "gholi-fly-agancy/internal/factor/port"
	reservationPort "gholi-fly-agancy/internal/reservation/port"
	staffPort "gholi-fly-agancy/internal/staff/port"
	tourPort "gholi-fly-agancy/internal/tour/port"
	tourEventPort "gholi-fly-agancy/internal/tour_event/port"

	"gorm.io/gorm"
)

type App interface {
	AgencyService(ctx context.Context) agencyPort.AgencyService
	StaffService(ctx context.Context) staffPort.StaffService
	FactorService(ctx context.Context) factorPort.FactorService
	TourService(ctx context.Context) tourPort.TourService
	TourEventService(ctx context.Context) tourEventPort.TourEventService
	ReservationService(ctx context.Context) reservationPort.ReservationService
	DB() *gorm.DB
	Config() config.Config
}
