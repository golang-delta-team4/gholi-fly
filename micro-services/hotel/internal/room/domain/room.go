package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	RoomID    = uint
	RoomUUID  = uuid.UUID
	RoomPrice = uint
)

type Room struct {
	ID          RoomID
	UUID        RoomUUID
	HotelID     uuid.UUID
	RoomNumber  uint
	Floor       uint
	BasePrice   RoomPrice
	AgencyPrice RoomPrice
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
