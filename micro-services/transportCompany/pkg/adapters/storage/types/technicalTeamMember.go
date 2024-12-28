package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TechnicalTeamMember struct {
	gorm.Model
	TechnicalTeamId uuid.UUID     `gorm:"not null;"`
	UserId          uuid.UUID     `gorm:"not null;"`
	Position        string        `gorm:"not null"`
	TechnicalTeam   TechnicalTeam `gorm:"foreignKey:TechnicalTeamId;constraint:OnDelete:CASCADE;"`
}
