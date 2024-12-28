package domain

import (
	"errors"
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

func (b *Booking) Validate() error {
	if b.HotelID == uuid.Nil {
		return errors.New("hotel id can't be nil")
	}
	if b.RoomID == uuid.Nil {
		return errors.New("room id can't be empty")
	}
	if b.ReservationID == uuid.Nil {
		return errors.New("reservation id can't be empty")
	}
	if b.CheckIn.IsZero() {
		return errors.New("check in can't be zero")
	}
	if b.CheckOut.IsZero() {
		return errors.New("check out can't be zero")
	}
	if b.CheckOut.Before(b.CheckIn) {
		return errors.New("check in can't be before check out")
	}
	if b.CheckOut.Before(time.Now()) {
		return errors.New("check in can't be in past")
	}
	if b.Status <= 0 {
		return errors.New("status not defined")
	}
	return nil
}
