package types

import (
	"time"

	"github.com/google/uuid"
)

type Factor struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	AgencyID          uuid.UUID `gorm:"type:uuid;index"` // References the Agency ID
	Agency            Agency    `gorm:"foreignKey:AgencyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HotelFactorID     uuid.UUID `gorm:"type:uuid;index"`
	TransportFactorID uuid.UUID `gorm:"type:uuid;index"`
	ReservationID     uuid.UUID `gorm:"type:uuid;index"`
	AgencyPrice       uint64
	Profit            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `gorm:"index"`
}
