package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;unique;not null;primaryKey"`
	HotelID     uuid.UUID `gorm:"type:uuid;references:UUID;not null"`
	Hotel       Hotel     `gorm:"foreignKey:HotelID;references:UUID;constraint:OnDelete:CASCADE"`
	RoomNumber  uint
	Floor       uint
	Size        uint
	BasePrice   uint
	AgencyPrice uint
}

func (h *Room) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
