package presenter

import (
	"time"

	"github.com/google/uuid"
)

type MatchMakerRequest struct {
	TripID             string   `json:"trip_id"`
	ReserveStartDate   string   `json:"reserve_start_date"`
	ReserveEndDate     string   `json:"reserve_end_date"`
	TripDistance       int      `json:"trip_distance"`
	NumberOfPassengers int      `json:"number_of_passengers"`
	TripType           VehicleType `json:"trip_type"`
	MaxPrice           int      `json:"max_price"`
	YearOfManufacture  int      `json:"year_of_manufacture"`
}

type VehicleType string

var (
	Bus      VehicleType = "bus"
	Train    VehicleType = "train"
	Ship     VehicleType = "ship"
	Airplane VehicleType = "airplane"
)

type MatchMakerResponse struct {
	ReservationID  uuid.UUID `json:"reservation_id"`
	VehicleDetails Vehicle   `json:"vehicle_detail"`
	Error          string    `json:"error"`
}

type Vehicle struct {
	ID                uuid.UUID `json:"id"`
	OwnerID           uuid.UUID `json:"owner_id"`
	Type              string    `json:"type"` // e.g., bus, train, plane, ship
	Capacity          int       `json:"capacity"`
	Speed             float64   `json:"speed"`
	UniqueCode        string    `json:"unique_code"`
	Status            string    `json:"status"` // active, inactive
	YearOfManufacture int       `json:"year_of_manufacture"`
	PricePerKilometer float64   `json:"price_per_kilometer"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
