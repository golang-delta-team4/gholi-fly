package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateHotel(ctx context.Context, hotel hotelDomain.Hotel) (hotelDomain.HotelUUID, error)
	GetAllHotels(ctx context.Context) ([]hotelDomain.Hotel, error)
	GetAllHotelsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]hotelDomain.Hotel, error)
	GetHotelByID(ctx context.Context, hotelID hotelDomain.HotelUUID) (*hotelDomain.Hotel, error)
	UpdateHotel(ctx context.Context, hotel hotelDomain.Hotel) error
	DeleteHotel(ctx context.Context, hotelID hotelDomain.HotelUUID) error
}
