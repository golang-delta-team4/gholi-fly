package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type StaffID uuid.UUID

type Staff struct {
	ID        StaffID
	AgencyID  uuid.UUID
	FirstName string
	LastName  string
	Role      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Staff) Validate() error {
	if s.AgencyID == uuid.Nil || s.FirstName == "" || s.LastName == "" || s.Role == "" || s.Email == "" {
		return errors.New("invalid staff details")
	}
	return nil
}
