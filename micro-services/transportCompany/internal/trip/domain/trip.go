package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	Id                       uuid.UUID
	CompanyID                uuid.UUID
	TripType                 string
	UserReleaseDate          time.Time
	TourReleaseDate          time.Time
	UserPrice                float64
	AgencyPrice              float64
	PathID                   uuid.UUID
	FromCountry              string
	ToCountry                string
	Origin                   string
	FromTerminalName         string
	ToTerminalName           string
	Destination              string
	PathName                 string
	PathDistanceKM           float64
	Status                   string
	MinPassengers            uint
	TechnicalTeamID          *uuid.UUID
	VehicleRequestID         *uuid.UUID
	VehicleYearOfManufacture int
	SoldTickets              uint
	MaxTickets               uint
	VehicleID                *uuid.UUID
	VehicleName              string
	IsCanceled               bool
	IsFinished               bool
	IsConfirmed              bool
	StartDate                *time.Time
	EndDate                  *time.Time
	Profit                   float64
	CreatedAt                time.Time
}

func (t *Trip) Validate() error {
	if t.CompanyID == uuid.Nil {
		return errors.New("company id can't be nil")
	}
	if t.TripType == "" {
		return errors.New("trip type can't be empty")
	}
	if !isValidVehicleType(t.TripType) {
		return errors.New("trip type must be one of 'bus', 'train', 'ship', or 'airplane'")
	}
	if t.UserReleaseDate.IsZero() {
		return errors.New("user release date can't be zero")
	}
	if t.TourReleaseDate.IsZero() {
		return errors.New("tourReleaseDate can't be zero")
	}
	if t.UserReleaseDate.Before(t.TourReleaseDate) {
		return errors.New("tour release date can't be before user release date")
	}
	if t.UserPrice < 0 {
		return errors.New("user price can't be negative")
	}
	if t.AgencyPrice < 0 {
		return errors.New("agency price can't be negative")
	}
	if t.PathID == uuid.Nil {
		return errors.New("path id can't be nil")
	}
	if t.MinPassengers == 0 {
		return errors.New("min passengers can't be zero")
	}
	if t.SoldTickets >= t.MaxTickets {
		return errors.New("sold tickets can't be greater than max tickets")
	}
	if t.Profit < 0 {
		return errors.New("profit can't be negative")
	}
	if t.StartDate.Before(time.Now()) {
		return errors.New("start date can't be before now")
	}
	if t.EndDate.Before(*t.StartDate) {
		return errors.New("end date can't be before start date")
	}
	if t.VehicleYearOfManufacture == 0 {
		return errors.New("vehicle year of manufacture can't be zero")
	}
	if t.AgencyPrice <= 0 {
		return errors.New("agency price can't be zero or negative")
	}
	if t.UserPrice <= 0 {
		return errors.New("user price can't be zero or negative")
	}
	return nil
}

type VehicleType string

var (
	Bus      VehicleType = "bus"
	Train    VehicleType = "train"
	Ship     VehicleType = "ship"
	Airplane VehicleType = "airplane"
)

func isValidVehicleType(tripType string) bool {
	switch VehicleType(tripType) {
	case Bus, Train, Ship, Airplane:
		return true
	default:
		return false
	}
}
