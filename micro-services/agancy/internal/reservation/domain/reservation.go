package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ReservationID uuid.UUID

type Reservation struct {
	ID         ReservationID
	CustomerID uuid.UUID
	FactorID   uuid.UUID
	TourID     uuid.UUID
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r *Reservation) Validate() error {
	if r.CustomerID == uuid.Nil || r.FactorID == uuid.Nil || r.TourID == uuid.Nil {
		return errors.New("invalid reservation details")
	}
	if r.Status == "" {
		return errors.New("status cannot be empty")
	}
	return nil
}
