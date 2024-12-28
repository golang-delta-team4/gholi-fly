package http

import (
	"context"

	"gholi-fly-agancy/api/service"
	"gholi-fly-agancy/app"
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
		return service.NewTourService(appContainer.TourService(ctx))
	}
}
