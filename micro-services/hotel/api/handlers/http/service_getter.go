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

func roomServiceGetter(appContainer app.App) ServiceGetter[*service.RoomService] {

	return func(ctx context.Context) *service.RoomService {
		return service.NewRoomService(appContainer.RoomService(ctx))
	}
}

func staffServiceGetter(appContainer app.App) ServiceGetter[*service.StaffService] {

	return func(ctx context.Context) *service.StaffService {
		return service.NewStaffService(appContainer.StaffService(ctx))
	}
}

func bookingServiceGetter(appContainer app.App) ServiceGetter[*service.BookingService] {

	return func(ctx context.Context) *service.BookingService {
		return service.NewBookingService(appContainer.BookingService(ctx))
	}
}
