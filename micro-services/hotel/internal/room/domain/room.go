package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	RoomID   = uint
	RoomUUID = uuid.UUID
)

type Room struct {
	ID          RoomID
	UUID        RoomUUID
	HotelID     uuid.UUID
	RoomNumber  uint
	Floor       uint
	BasePrice   uint
	AgencyPrice uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
