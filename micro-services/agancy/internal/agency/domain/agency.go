package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AgencyID uuid.UUID

type Agency struct {
	ID               AgencyID
	Name             string
	OwnerID          uuid.UUID
	WalletID         uuid.UUID
	ProfitPercentage float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (a *Agency) Validate() error {
	if a.Name == "" || a.ProfitPercentage < 0 {
		return errors.New("invalid agency details")
	}
	return nil
}
