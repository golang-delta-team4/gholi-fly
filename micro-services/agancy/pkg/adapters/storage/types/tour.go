package types

import (
	"time"

	"github.com/google/uuid"
)

type Tour struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name                string
	Description         string
	StartDate           time.Time
	EndDate             time.Time
	SourceLocation      string
	DestinationLocation string
	ForwardTripID       uuid.UUID `gorm:"type:uuid;index"` // References a Trip
	BackwardTripID      uuid.UUID `gorm:"type:uuid;index"` // References a Trip
	TripAgencyPrice     int
	HotelID             uuid.UUID `gorm:"type:uuid;index"` // References a Hotel
	HotelAgencyPrice    int
	IsPublished         bool
	Capacity            int
	Reservations        []Reservation `gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-many with Reservation
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `gorm:"index"`
}
