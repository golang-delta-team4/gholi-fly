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
	ProductionYearMin     int
	Status                string    `gorm:"default:'pending'"`
	MatchedVehicleID      uuid.UUID `gorm:"type:uuid;not null;"`
	VehicleReservationFee float64
	VehicleProductionYear int
	VehicleName           string
	MatchVehicleSpeed     float64
	MinCost               float64
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}

func (base *VehicleRequest) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == uuid.Nil {
		base.Id = uuid.New()
	}
	return
}
