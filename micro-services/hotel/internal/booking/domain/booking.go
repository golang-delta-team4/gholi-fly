package domain

import (
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	"time"

	"github.com/google/uuid"
)

type (
	BookingID   = uint
	BookingUUID = uuid.UUID
)

type BookingStatus = uint8

const (
	BookingStatusUnknown BookingStatus = iota
	BookingStatusPending
	BookingStatusConfirmed
	BookingStatusCancelled
)

type Booking struct {
	ID            BookingID
	UUID          BookingUUID
	HotelID       hotelDomain.HotelUUID
	RoomID        uuid.UUID
	UserID        *uuid.UUID
	AgencyID      *uuid.UUID
	ReservationID uuid.UUID
	CheckIn       time.Time
	CheckOut      time.Time
	Status        BookingStatus
	IsPayed       bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
