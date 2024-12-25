package domain

import (
	"time"
	"github.com/google/uuid"
)

// Terminal represents a terminal (bus station, airport, etc.)
type Terminal struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
