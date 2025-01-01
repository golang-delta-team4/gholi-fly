package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	UUID       uuid.UUID
	Route      string
	Method     string
	ResourceID *uint
	CreatedAt time.Time
}
