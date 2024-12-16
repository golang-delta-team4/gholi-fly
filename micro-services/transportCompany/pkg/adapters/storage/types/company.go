package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string
	Description string
	OwnerId     uuid.UUID `gorm:"type:uuid;not null; uniqueIndex"`
}
