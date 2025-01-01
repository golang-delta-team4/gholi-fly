package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID       uuid.UUID
	FirstName  string
	LastName   string
	Email      string `gorm:"unique"`
	Password   string
	IsBlocked bool `gorm:"default:false"`
	IsVerified bool
	UserRoles []UserRole
}

type RefreshToken struct {
	gorm.Model
	UserID         uint
	User           *User `gorm:"foreignKey:UserID"`
	Token          string
	ExpirationTime time.Time
}

type UserAuthorization struct {
	UserUUID uuid.UUID
	Route    string
	Method   string
}
