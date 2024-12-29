package domain

import (
	"errors"
	"time"

	roomDomain "gholi-fly-hotel/internal/room/domain"

	"github.com/google/uuid"
)

type (
	HotelID   = uint
	HotelUUID = uuid.UUID
	OwnerUUID = uuid.UUID
)

func HotelUUIDFromString(s string) (HotelUUID, error) {
	uid, err := uuid.Parse(s)
	return HotelUUID(uid), err
}

type Hotel struct {
	UUID      HotelUUID
	OwnerID   OwnerUUID
	Name      string
	City      string
	Rooms     []roomDomain.Room
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type HotelFilters struct {
	Name string
	City string
}

func (h *Hotel) Validate() error {
	if h.Name == "" {
		return errors.New("hotel name cant be empty")
	}
	if h.City == "" {
		return errors.New("city cant cant be empty")
	}
	if h.OwnerID == uuid.Nil {
		return errors.New("owner id cant be nil")
	}

	return nil
}
