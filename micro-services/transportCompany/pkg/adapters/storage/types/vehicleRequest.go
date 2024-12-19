package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleRequest struct {
	Id                    uuid.UUID `gorm:"type:uuid;primaryKey"`
	TripID                uuid.UUID `gorm:"type:uuid;not null;"`
	VehicleType           string    `gorm:"not null"`
	MinCapacity           int
	Status                string    `gorm:"default:'pending'"`
	MatchedVehicleID      uuid.UUID `gorm:"type:uuid;not null;"`
	VehicleProductionYear int
	VehicleName           string
	Cost                  float64
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
