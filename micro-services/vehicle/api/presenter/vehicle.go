package presenter

import (
	"vehicle/internal/vehicle/domain"

	"github.com/google/uuid"
)

type Vehicle struct {
	Capacity          int         `json:"capacity"`
	Speed             float64     `json:"speed"`
	Status            Status      `json:"status"`
	PricePerKilometer float64     `json:"price_per_kilometer"`
}

type MatchMakerRequest struct {
	TripID             uuid.UUID
	ReserveStartDate   string
	ReserveEndDate     string
	TripDistance       int
	NumberOfPassengers int
	TripType           domain.VehicleType
	MaxPrice           int
	YearOfManufacture  int
}

type Status string

var (
	Active Status = "active"
	InActive Status = "inactive"
)