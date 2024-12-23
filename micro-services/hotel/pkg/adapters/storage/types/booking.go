package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"unique"`
	HotelID  uint
	Hotel    Hotel `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE"`
	RoomID   uint
	Room     Room `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
	UserID   *uuid.UUID
	AgencyID *uuid.UUID
	CheckIn  time.Time
	CheckOut time.Time
	Status   string
}

func (h *Booking) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
