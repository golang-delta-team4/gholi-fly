package types

import (
	"gholi-fly-agancy/internal/tour_event/domain"
	"time"

	"github.com/google/uuid"
)

type TourEvent struct {
	ID                  uuid.UUID          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ReservationID       uuid.UUID          `gorm:"type:uuid;index"` // References a Reservation
	EventType           domain.EventType   `gorm:"type:varchar(50);index"`
	Payload             domain.JSONB       `gorm:"type:jsonb"` // JSON payload for event data
	CompensationPayload domain.JSONB       `gorm:"type:jsonb"` // JSON payload for compensation data
	Status              domain.EventStatus `gorm:"type:varchar(50);index"`
	RetryCount          int                `gorm:"default:0"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `gorm:"index"`
}
