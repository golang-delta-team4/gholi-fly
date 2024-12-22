package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID        uuid.UUID
	Name        string       `gorm:"unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	User      *User `gorm:"foreignKey:UserID"`
	RoleID    uint
	Role      *Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
}
