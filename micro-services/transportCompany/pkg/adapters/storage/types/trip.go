package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trip struct {
	Id               uuid.UUID `gorm:"type:uuid;primaryKey"`
	CompanyID        uuid.UUID `gorm:"not null"`
	Company          *Company  `gorm:"foreignKey:CompanyID; constraint:OnDelete:CASCADE;"`
	TripType         string    `gorm:"not null"`
	UserReleaseDate  time.Time
	TourReleaseDate  time.Time
	UserPrice        float64
	AgencyPrice      float64
	PathID           uuid.UUID `gorm:"not null"`
	FromCountry      string
	ToCountry        string
	Origin           string `gorm:"not null"`
	FromTerminalName string
	ToTerminalName   string
	Destination      string `gorm:"not null"`
	PathName         string
	PathDistanceKM   float64
	Status           string `gorm:"default:'pending'"`
	MinPassengers    uint
	TechnicalTeamID  *uuid.UUID     `gorm:"type:uuid;null"`
	TechnicalTeam    *TechnicalTeam `gorm:"foreignKey:TechnicalTeamID; constraint:OnDelete:CASCADE;"`
	VehicleRequestID *uuid.UUID
	VehicleRequest   *VehicleRequest `gorm:"foreignKey:TripID; constraint:OnDelete:CASCADE;"`
	Tickets          []Ticket        `gorm:"foreignKey:TripID; constraint:OnDelete:CASCADE;"`
	SoldTickets      uint            `gorm:"default:0"`
	MaxTickets       uint
	VehicleID        *uuid.UUID
	VehicleName      string
	IsCanceled       bool       `gorm:"default:false"`
	IsFinished       bool       `gorm:"default:false"`
	IsConfirmed      bool       `gorm:"default:false"`
	StartDate        *time.Time `gorm:"not null;"`
	EndDate          *time.Time
	Profit           float64 `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (base *Trip) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == uuid.Nil {
		base.Id = uuid.New()
	}
	return
}
