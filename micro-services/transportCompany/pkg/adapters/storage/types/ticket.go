package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	Id        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	TripID    uuid.UUID  `gorm:"not null"`
	Trip      *Trip      `gorm:"foreignKey:TripID; constraint:OnDelete:CASCADE;"`
	UserID    *uuid.UUID `gorm:"type:uuid;default:NULL"`
	AgencyID  *uuid.UUID `gorm:"type:uuid;default:NULL"`
	Price     float64
	Status    string `gorm:"default:'pending'"`
	InvoiceId uuid.UUID
	Invoice   Invoice `gorm:"foreignKey:InvoiceId; constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == uuid.Nil {
		base.Id = uuid.New()
	}
	return
}
