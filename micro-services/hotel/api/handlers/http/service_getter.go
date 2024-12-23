package http

import (
	"context"

	"gholi-fly-hotel/api/service"
	"gholi-fly-hotel/app"
)

// hotel service transient instance handler
func hotelServiceGetter(appContainer app.App) ServiceGetter[*service.HotelService] {

	return func(ctx context.Context) *service.HotelService {
		return service.NewHotelService(appContainer.HotelService(ctx))
	}
}
