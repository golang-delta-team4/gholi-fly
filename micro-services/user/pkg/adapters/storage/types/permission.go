package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	UUID       uuid.UUID
	Route      string
	Method     string
}
