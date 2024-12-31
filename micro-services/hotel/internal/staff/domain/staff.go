package domain

import (
	"errors"
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
	ID      StaffID
	UUID    StaffUUID
	HotelID hotelDomain.HotelUUID
	// WalletID
	Name      string
	StaffType StaffType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (s *Staff) Validate() error {
	if s.HotelID == uuid.Nil {
		return errors.New("hotel id can't be nil")
	}
	if s.StaffType <= 0 {
		return errors.New("undefined staff type")
	}
	if s.Name == "" {
		return errors.New("name cant be empty")
	}

	return nil
}
