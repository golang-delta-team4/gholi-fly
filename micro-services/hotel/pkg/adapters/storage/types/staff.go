package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"unique;primaryKey"`
	HotelID   uuid.UUID `gorm:"type:uuid;references:UUID;not null;uniqueIndex:idx_hotel_name"`
	Hotel     Hotel     `gorm:"foreignKey:HotelID;references:UUID;constraint:OnDelete:CASCADE"`
	Name      string    `gorm:"unique;not null;required;uniqueIndex:idx_hotel_name"`
	StaffType uint8     `gorm:"not null;default:0"`
}

func (h *Staff) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
