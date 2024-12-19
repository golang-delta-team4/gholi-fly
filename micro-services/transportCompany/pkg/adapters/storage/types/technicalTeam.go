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
	CompanyId   uint    `gorm:"not null;unique;"`
	TripType    string  `gorm:"not null"`
	Company     Company `gorm:"foreignKey:CompanyId;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
