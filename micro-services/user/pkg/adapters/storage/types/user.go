package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
}
