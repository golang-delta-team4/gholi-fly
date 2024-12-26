package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ReservationID uuid.UUID

type Reservation struct {
	ID         ReservationID
	ToureID    uuid.UUID
	UserID     uuid.UUID
	Seats      int
	TotalPrice float64
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r *Reservation) Validate() error {
	if r.ToureID == uuid.Nil || r.UserID == uuid.Nil || r.Seats <= 0 || r.TotalPrice < 0 {
		return errors.New("invalid reservation details")
	}
	if r.Status == "" {
		return errors.New("status cannot be empty")
	}
	return nil
}
