package types

import (
	"gorm.io/gorm"
)

type TechnicalTeam struct {
	gorm.Model
	Name        string
	Description string
	CompanyId   uint    `gorm:"not null;unique;"`
	Company     Company `gorm:"foreignKey:CompanyId;constraint:OnDelete:CASCADE;"`
}
