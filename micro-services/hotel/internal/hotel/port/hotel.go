package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
)

type Repo interface {
	Create(ctx context.Context, hotel hotelDomain.Hotel) (hotelDomain.HotelUUID, error)
	GetByID(ctx context.Context, hotelID hotelDomain.HotelUUID) (*hotelDomain.Hotel, error)
	GetAll(ctx context.Context) ([]hotelDomain.Hotel, error)
	Update(ctx context.Context, hotel hotelDomain.Hotel) error
	Delete(ctx context.Context, hotelID hotelDomain.HotelUUID) error
}
