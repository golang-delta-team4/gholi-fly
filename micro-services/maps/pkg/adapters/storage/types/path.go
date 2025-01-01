package types

import (
	"time"

	"github.com/google/uuid"
)

// Path is the database representation of a transportation path.
type Path struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SourceTerminalID      uuid.UUID `gorm:"type:uuid;not null"`
	DestinationTerminalID uuid.UUID `gorm:"type:uuid;not null"`
	DestinationTerminal   *Terminal `gorm:"foreignKey:DestinationTerminalID"`
	SourceTerminal        *Terminal `gorm:"foreignKey:SourceTerminalID"`
	DistanceKM            float64   `gorm:"type:float;not null"`
	RouteCode             string    `gorm:"type:varchar(100);unique;not null"`
	VehicleType           string    `gorm:"type:varchar(50);not null"`
	CreatedAt             time.Time `gorm:"autoCreateTime"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime"`
}
