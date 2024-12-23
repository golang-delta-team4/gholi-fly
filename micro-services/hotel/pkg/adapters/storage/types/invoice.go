package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"unique"`
	FactorId  uuid.UUID
	BookingID uuid.UUID
	Booking   Booking `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	PaidAt    *time.Time
	IsPaid    bool
}

func (h *Invoice) BeforeCreate(tx *gorm.DB) error {
	if h.UUID == uuid.Nil {
		h.UUID = uuid.New()
	}
	return nil
}
