package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type StaffID uuid.UUID

type Staff struct {
	ID        StaffID
	UserID    uuid.UUID
	AgencyID  uuid.UUID
	WalletID  uuid.UUID
	Stock     int
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Staff) Validate() error {
	if s.UserID == uuid.Nil || s.AgencyID == uuid.Nil || s.Role == "" {
		return errors.New("invalid staff details")
	}
	return nil
}
