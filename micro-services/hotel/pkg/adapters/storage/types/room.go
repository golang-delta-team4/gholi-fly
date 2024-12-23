package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"unique"`
	HotelID     uuid.UUID
	Hotel       Hotel `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE"`
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
