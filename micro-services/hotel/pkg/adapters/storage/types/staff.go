package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"unique"`
	HotelID   uuid.UUID
	Hotel     Hotel `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE"`
	Name      string
	StaffType string
}

func (h *Staff) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
