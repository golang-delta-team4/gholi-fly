package presenter

import (
	"github.com/google/uuid"
)

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
