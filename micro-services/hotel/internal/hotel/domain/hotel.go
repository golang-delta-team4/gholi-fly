package domain

import (
	"time"

	"github.com/google/uuid"
)

type HotelUUID = uuid.UUID

type Hotel struct {
	ID        HotelUUID
	OwnerID   uuid.UUID
	Name      string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
