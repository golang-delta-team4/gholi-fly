package types

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID uuid.UUID `gorm:"type:uuid;index"`
	FactorID   uuid.UUID `gorm:"type:uuid;index"` // References Factor ID
	Factor     Factor    `gorm:"foreignKey:FactorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TourID     uuid.UUID `gorm:"type:uuid;index"` // References Tour ID
	Tour       Tour      `gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}
