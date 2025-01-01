package domain

import (
	"errors"
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

func (r *Room) Validate() error {
	if r.HotelID == uuid.Nil {
		return errors.New("hotel id can't be nil")
	}
	if r.RoomNumber <= 0 {
		return errors.New("room number cant be zero or negative")
	}
	if r.Floor <= 0 {
		return errors.New("floor cant be zero or negative")
	}
	if r.BasePrice <= 0 {
		return errors.New("base price cant be zero or negative")
	}
	if r.AgencyPrice <= 0 {
		return errors.New("agency price cant be zero or negative")
	}

	return nil
}
