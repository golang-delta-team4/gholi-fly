package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	HotelID   = uint
	HotelUUID = uuid.UUID
)

func HotelUUIDFromString(s string) (HotelUUID, error) {
	uid, err := uuid.Parse(s)
	return HotelUUID(uid), err
}

type Hotel struct {
	UUID      HotelUUID
	OwnerID   uuid.UUID
	Name      string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type HotelFilters struct {
	Name string
	City string
}
