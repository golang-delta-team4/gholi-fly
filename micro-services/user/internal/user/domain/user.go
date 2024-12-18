package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
