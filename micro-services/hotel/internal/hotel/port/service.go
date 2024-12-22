package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
)

type Service interface {
	CreateHotel(ctx context.Context, hotel hotelDomain.Hotel) (hotelDomain.HotelUUID, error)
	GetHotelByID(ctx context.Context, hotelID hotelDomain.HotelUUID) (*hotelDomain.Hotel, error)
	GetHotels(ctx context.Context) ([]hotelDomain.Hotel, error)
	UpdateHotel(ctx context.Context, hotel hotelDomain.Hotel) error
	DeleteHotel(ctx context.Context, hotelID hotelDomain.HotelUUID) error
}
