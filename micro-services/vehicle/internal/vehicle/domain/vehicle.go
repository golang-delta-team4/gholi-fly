package domain

import (
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OwnerID           uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	Type              string    `gorm:"type:varchar(50);not null" json:"type"` // e.g., bus, train, plane, ship
	Capacity          int       `gorm:"not null" json:"capacity"`
	Speed             float64   `gorm:"not null" json:"speed"`
	UniqueCode        string    `gorm:"type:varchar(100);unique;not null" json:"unique_code"`
	Status            string    `gorm:"type:varchar(20);not null" json:"status"` // active, inactive
	YearOfManufacture int       `gorm:"type:int;not null" json:"year_of_manufacture"`
	PricePerKilometer float64
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MatchMakerRequest struct {
	TripID             uuid.UUID
	ReserveStartDate   time.Time
	ReserveEndDate     time.Time
	TripDistance       int
	NumberOfPassengers int
	TripType           VehicleType
	MaxPrice           int
	YearOfManufacture  int
}

type VehicleType string

var (
	Bus   VehicleType = "bus"
	Train VehicleType = "train"
	Ship  VehicleType = "ship"
	Airplane VehicleType = "airplane"
)