package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TechnicalTeam struct {
	Id          uuid.UUID
	Name        string
	Description string
	CompanyId   uuid.UUID
	TripType    string
	Members     []string
}

func (t *TechnicalTeam) Validate() error {

	if t.Name == "" {
		return errors.New("name is required")
	}

	if t.CompanyId == uuid.Nil {
		return errors.New("company id is required")
	}

	if t.TripType == "" {
		return errors.New("trip type is required")
	}
	return nil
}

type TechnicalTeamMember struct {
	UserId   uuid.UUID
	TeamId   uuid.UUID
	Position string
}
