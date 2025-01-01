package types

import (
	"time"

	"github.com/google/uuid"
)

// relation btn terminal and path here!!


// TerminalType defines the possible types of terminals
type TerminalType string

const (
	BusStation   TerminalType = "bus_station"
	Airport      TerminalType = "airport"
	TrainStation TerminalType = "train_station"
	SeaPort      TerminalType = "sea_port"
)

// Terminal is the database representation of a terminal
type Terminal struct {
	ID        uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string      `gorm:"type:varchar(255);not null"`
	Location  string      `gorm:"type:varchar(255);not null"`
	Type      TerminalType `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}
