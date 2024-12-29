package types

import (
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OwnerID           uuid.UUID `gorm:"type:uuid;not null"`
	Type              string    `gorm:"type:varchar(50);not null"` // e.g., bus, train, plane, ship
	Capacity          int       `gorm:"not null"`
	Speed             float64   `gorm:"not null"`
	UniqueCode        string    `gorm:"type:varchar(100);unique;not null"`
	Status            string    `gorm:"type:varchar(20);not null"` // active, inactive
	YearOfManufacture int       `gorm:"type:int;not null"`         // Year vehicle was manufactured
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
