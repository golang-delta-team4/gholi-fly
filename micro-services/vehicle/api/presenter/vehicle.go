package presenter

import (
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID                uuid.UUID   `json:"id"`
	OwnerID           uuid.UUID   `json:"owner_id"`
	Type              VehicleType `json:"type"` // e.g., bus, train, plane, ship
	Capacity          int         `json:"capacity"`
	Speed             float64     `json:"speed"`
	UniqueCode        string      `json:"unique_code"`
	Status            string      `json:"status"`
	YearOfManufacture int         `json:"year_of_manufacture"`
	PricePerKilometer float64     `json:"price_per_kilometer"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
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
