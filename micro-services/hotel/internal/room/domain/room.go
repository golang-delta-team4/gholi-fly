package domain

import (
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	"time"

	"github.com/google/uuid"
)

type RoomUUID = uuid.UUID

type Room struct {
	ID          RoomUUID
	HotelID     hotelDomain.HotelUUID
	RoomNumber  uint
	Floor       uint
	BasePrice   uint
	AgencyPrice uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
