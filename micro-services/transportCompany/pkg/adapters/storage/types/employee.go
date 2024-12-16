package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	UserId    uuid.UUID `gorm:"type:uuid;not null; uniqueIndex"`
	CompanyId uint
	Company   Company `gorm:"foreignKey:CompanyId"`
}
