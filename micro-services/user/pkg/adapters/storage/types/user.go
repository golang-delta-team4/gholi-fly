package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID         uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	Password     string
	IsVerified   bool
	Roles []Role `gorm:"many2many:user_roles;"`
}

type RefreshToken struct {
	gorm.Model
	UserID         uint
	User           *User `gorm:"foreignKey:UserID"`
	Token          string
	ExpirationTime time.Time
}
