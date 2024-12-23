package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"unique"`
	OwnerEmail string
	Name       string `gorm:"unique"`
	City       string
}

func (h *Hotel) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
