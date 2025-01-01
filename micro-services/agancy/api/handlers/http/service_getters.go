package http

import (
	"context"

	"gholi-fly-agancy/api/service"
	"gholi-fly-agancy/app"
	"gholi-fly-agancy/pkg/adapters/clients/grpc"
)

// agencyService transient instance handler
func agencyServiceGetter(appContainer app.App) ServiceGetter[*service.AgencyService] {
	return func(ctx context.Context) *service.AgencyService {
		return service.NewAgencyService(appContainer.AgencyService(ctx), appContainer.StaffService(ctx))
	}
}

// tourService transient instance handler
func tourServiceGetter(appContainer app.App) ServiceGetter[*service.TourService] {
	return func(ctx context.Context) *service.TourService {
		return service.NewTourService(appContainer.TourService(ctx), appContainer.TourEventService(ctx), appContainer.AgencyService(ctx))
	}
}

// reservationService transient instance handler
func reservationServiceGetter(appContainer app.App) ServiceGetter[*service.ReservationService] {
	return func(ctx context.Context) *service.ReservationService {
		return service.NewReservationService(
			appContainer.TourService(ctx),
			appContainer.ReservationService(ctx),
			appContainer.FactorService(ctx),
			appContainer.AgencyService(ctx),
			grpc.NewGRPCBankClient(appContainer.Config().Bank.Host, appContainer.Config().Bank.Port),
		)
	}
}
