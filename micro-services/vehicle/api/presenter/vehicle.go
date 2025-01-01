package presenter

import (
	"vehicle/internal/vehicle/domain"

)

type Vehicle struct {
	Capacity          int         `json:"capacity"`
	Speed             float64     `json:"speed"`
	Status            Status      `json:"status"`
	PricePerKilometer float64     `json:"price_per_kilometer"`
}

type MatchMakerRequest struct {
	TripID             string   `json:"trip_id"`
	ReserveStartDate   string   `json:"reserve_start_date"`
	ReserveEndDate     string   `json:"reserve_end_date"`
	TripDistance       int      `json:"trip_distance"`
	NumberOfPassengers int      `json:"number_of_passengers"`
	TripType           domain.VehicleType `json:"trip_type"`
	MaxPrice           int      `json:"max_price"`
	YearOfManufacture  int      `json:"year_of_manufacture"`
}


type Status string

var (
	Active Status = "active"
	InActive Status = "inactive"
)