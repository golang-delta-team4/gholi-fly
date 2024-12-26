package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UUID     uuid.UUID   `gorm:"unique;primaryKey"`
	HotelID  uuid.UUID   `gorm:"type:uuid;not null;references:UUID"`
	Hotel    Hotel       `gorm:"foreignKey:HotelID;references:UUID;constraint:OnDelete:CASCADE"`
	RoomIDs  []uuid.UUID `gorm:"type:uuid[]"`
	UserID   *uuid.UUID
	AgencyID *uuid.UUID
	CheckIn  time.Time
	CheckOut time.Time
	Status   uint8
}

func (h *Booking) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
