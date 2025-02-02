package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UUID          uuid.UUID `gorm:"unique;primaryKey"`
	HotelID       uuid.UUID `gorm:"type:uuid;not null;references:UUID"`
	Hotel         Hotel     `gorm:"foreignKey:HotelID;references:UUID;constraint:OnDelete:CASCADE"`
	RoomID        uuid.UUID `gorm:"type:uuid;not null;references:UUID"`
	Room          Room      `gorm:"foreignKey:RoomID;references:UUID;constraint:OnDelete:CASCADE"`
	UserID        uuid.UUID
	FactorID      string
	ReservationID uuid.UUID
	IsPaid        bool
	PaidDate      *time.Time
	CheckIn       time.Time
	CheckOut      time.Time
	Status        uint8
}

func (h *Booking) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
