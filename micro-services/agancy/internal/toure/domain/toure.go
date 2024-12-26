package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ToureID uuid.UUID

type Toure struct {
	ID          ToureID
	Name        string
	Description string
	AgencyID    uuid.UUID
	Price       float64
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Toure) Validate() error {
	if t.Name == "" || t.Price < 0 || t.StartDate.After(t.EndDate) || t.AgencyID == uuid.Nil {
		return errors.New("invalid toure details")
	}
	return nil
}
