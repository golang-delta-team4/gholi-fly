package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CompanyId uuid.UUID `gorm:"type:uuid;not null"`
	Position  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Company   Company        `gorm:"foreignKey:CompanyId;constraint:OnDelete:CASCADE;"`
}
