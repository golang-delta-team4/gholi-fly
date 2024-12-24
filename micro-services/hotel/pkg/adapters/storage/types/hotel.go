package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	UUID    uuid.UUID `gorm:"type:uuid;primaryKey;unique;not null"`
	OwnerID uuid.UUID `gorm:"type:uuid;not null;required"`
	Name    string    `gorm:"unique;not null;required"`
	City    string    `gorm:"not null;required"`
}

func (h *Hotel) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
