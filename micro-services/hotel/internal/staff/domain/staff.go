package domain

import (
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"

	"github.com/google/uuid"
)

type StaffUUID = uuid.UUID

type StaffType uint8

const (
	StaffTypeUnknown StaffType = iota
	StaffTypeManager
	StaffTypeReceptionist
	StaffTypeCleaner
	StaffTypeSecurity
)

type Staff struct {
	ID        StaffUUID
	HotelID   hotelDomain.HotelUUID
	Name      string
	StaffType StaffType
}
