package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	Id               uuid.UUID
	CompanyID        uuid.UUID
	TripType         string
	UserReleaseDate  time.Time
	TourReleaseDate  time.Time
	UserPrice        float64
	AgencyPrice      float64
	PathID           uuid.UUID
	FromCountry      string
	ToCountry        string
	Origin           string
	FromTerminalName string
	ToTerminalName   string
	Destination      string
	PathName         string
	PathDistanceKM   float64
	Status           string
	MinPassengers    uint
	TechnicalTeamID  *uuid.UUID
	VehicleRequestID *uuid.UUID
	SoldTickets      uint
	MaxTickets       uint
	VehicleID        *uuid.UUID
	VehicleName      string
	IsCanceled       bool
	IsFinished       bool
	IsConfirmed      bool
	StartDate        *time.Time
	EndDate          *time.Time
	Profit           float64
	CreatedAt        time.Time
}

func (t *Trip) Validate() error {
	if t.CompanyID == uuid.Nil {
		return errors.New("CompanyID can't be nil")
	}
	if t.TripType == "" {
		return errors.New("TripType can't be empty")
	}
	if t.UserReleaseDate.IsZero() {
		return errors.New("UserReleaseDate can't be zero")
	}
	if t.TourReleaseDate.IsZero() {
		return errors.New("TourReleaseDate can't be zero")
	}
	if t.UserReleaseDate.Before(t.TourReleaseDate) {
		return errors.New("Tour release date can't be before user release date")
	}
	if t.UserPrice < 0 {
		return errors.New("UserPrice can't be negative")
	}
	if t.AgencyPrice < 0 {
		return errors.New("AgencyPrice can't be negative")
	}
	if t.PathID == uuid.Nil {
		return errors.New("PathID can't be nil")
	}
	if t.MinPassengers == 0 {
		return errors.New("MinPassengers can't be zero")
	}
	if t.SoldTickets > t.MaxTickets {
		return errors.New("SoldTickets can't be greater than MaxTickets")
	}
	if t.Profit < 0 {
		return errors.New("Profit can't be negative")
	}
	if t.StartDate.Before(time.Now()) {
		return errors.New("StartDate can't be before now")
	}
	return nil
}
