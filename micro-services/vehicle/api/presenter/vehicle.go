package presenter

import (

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
	TripType           TripType
	MaxPrice           int
	YearOfManufacture  int
}

type TripType string

var (
	GroundTrip TripType = "ground"
	AirTrip    TripType = "air"
	SeaTrip    TripType = "sea"
)

type VehicleType string

var (
	Bus   VehicleType = "bus"
	Train VehicleType = "train"
	Ship  VehicleType = "ship"
	Plane VehicleType = "plane"
)

type Status string

var (
	Active Status = "active"
	InActive Status = "inactive"
)