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
	Email        string `gorm:"unique"`
	Password     string
	IsVerified   bool
}

type RefreshToken struct {
	gorm.Model
	UserID         uint
	User           *User `gorm:"foreignKey:UserID"`
	Token          string
	ExpirationTime time.Time
}
