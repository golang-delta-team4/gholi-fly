package domain

import (
	"time"
	"github.com/google/uuid"
)

type Path struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SourceTerminalID      uuid.UUID `gorm:"type:uuid;not null" json:"source_terminal_id"`
	DestinationTerminalID uuid.UUID `gorm:"type:uuid;not null" json:"destination_terminal_id"`
	DistanceKM            float64   `gorm:"type:float;not null" json:"distance_km"`
	RouteCode             string    `gorm:"type:varchar(100);unique;not null" json:"route_code"`
	VehicleType           string    `gorm:"type:varchar(50);not null" json:"vehicle_type"`
	CreatedAt             time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
