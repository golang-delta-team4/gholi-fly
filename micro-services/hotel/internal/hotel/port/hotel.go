package port

import (
	"context"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, hotel hotelDomain.Hotel) (hotelDomain.HotelUUID, error)
	Get(ctx context.Context) ([]hotelDomain.Hotel, error)
	GetByOwnerID(ctx context.Context, ownerId uuid.UUID) ([]hotelDomain.Hotel, error)
	GetByID(ctx context.Context, hotelID hotelDomain.HotelUUID) (*hotelDomain.Hotel, error)
	Update(ctx context.Context, hotel hotelDomain.Hotel) error
	Delete(ctx context.Context, hotelID hotelDomain.HotelUUID) error
}
