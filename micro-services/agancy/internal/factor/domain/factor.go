package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type FactorID uuid.UUID

type Factor struct {
	ID                FactorID
	HotelFactorID     uuid.UUID
	TransportFactorID uuid.UUID
	ReservationID     uuid.UUID
	AgencyPrice       uint64
	Profit            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (f *Factor) Validate() error {
	if f.HotelFactorID == uuid.Nil || f.TransportFactorID == uuid.Nil || f.ReservationID == uuid.Nil {
		return errors.New("invalid factor details")
	}
	if f.AgencyPrice < 0 || f.Profit < 0 {
		return errors.New("invalid pricing or profit details")
	}
	return nil
}
