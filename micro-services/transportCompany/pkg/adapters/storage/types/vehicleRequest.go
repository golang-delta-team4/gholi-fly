package types

import "gorm.io/gorm"

type VehicleRequest struct {
	gorm.Model
	TripID                uint
	VehicleType           string `gorm:"type:varchar(50);not null"`
	MinCapacity           int
	ProductionYearMin     int
	Status                string `gorm:"type:varchar(20);default:'pending'"`
	MatchedVehicleID      uint
	VehicleReservationFee float64
	VehicleProductionYear int
	VehicleName           string
	MatchVehicleSpeed     float64
	MinCost               float64
}
