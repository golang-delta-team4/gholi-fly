package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID uuid.UUID
	Name string `gorm:"unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}