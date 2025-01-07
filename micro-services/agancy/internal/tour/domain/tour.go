package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TourID uuid.UUID

func (id TourID) String() string {
	return uuid.UUID(id).String()
}

type Tour struct {
	ID                  TourID
	Name                string
	Description         string
	StartDate           time.Time
	EndDate             time.Time
	SourceLocation      string
	DestinationLocation string
	ForwardTripID       uuid.UUID
	BackwardTripID      uuid.UUID
	TripAgencyPrice     int
	HotelID             uuid.UUID
	HotelAgencyPrice    int
	IsPublished         bool
	Capacity            int
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
	if t.Capacity <= 0 {
		return errors.New("capacity must be greater than zero")
	}
	if t.TripAgencyPrice < 0 || t.HotelAgencyPrice < 0 {
		return errors.New("prices cannot be negative")
	}
	return nil
}
