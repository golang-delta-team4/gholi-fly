package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TechnicalTeam struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name        string
	Description string
	CompanyId   uuid.UUID `gorm:"not null;"`
	TripType    string    `gorm:"not null"`
	Company     Company   `gorm:"foreignKey:CompanyId;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (base *TechnicalTeam) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == uuid.Nil {
		base.Id = uuid.New()
	}
	return
}
