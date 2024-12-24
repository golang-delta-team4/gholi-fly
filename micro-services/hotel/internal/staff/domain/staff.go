package domain

import (
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	"time"

	"github.com/google/uuid"
)

type (
	StaffID   = uint
	StaffUUID = uuid.UUID
)

type StaffType = uint8

const (
	StaffTypeUnknown StaffType = iota
	StaffTypeOwner
	StaffTypeManager
	StaffTypeReceptionist
	StaffTypeCleaner
	StaffTypeSecurity
)

type Staff struct {
	ID        StaffID
	UUID      StaffUUID
	HotelID   hotelDomain.HotelUUID
	Name      string
	StaffType StaffType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
