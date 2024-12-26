package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TourID uuid.UUID

type Tour struct {
	ID                  TourID
	Name                string
	Description         string
	StartDate           time.Time
	EndDate             time.Time
	SourceLocation      string
	DestinationLocation string
	ForwardTicketID     uuid.UUID
	ReturnTicketID      uuid.UUID
	HotelBookingID      uuid.UUID
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (t *Tour) Validate() error {
	if t.Name == "" || t.StartDate.After(t.EndDate) {
		return errors.New("invalid tour details")
	}
	if t.SourceLocation == "" || t.DestinationLocation == "" {
		return errors.New("source and destination locations are required")
	}
	return nil
}
