package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TechnicalTeamMemeber struct {
	gorm.Model
	TechnicalTeamId int64         `gorm:"not null;unique;"`
	UserId          uuid.UUID     `gorm:"type:uuid;not null;unique;"`
	Position        string        `gorm:"not null"`
	TechnicalTeam   TechnicalTeam `gorm:"foreignKey:TechnicalTeamId;constraint:OnDelete:CASCADE;"`
}
