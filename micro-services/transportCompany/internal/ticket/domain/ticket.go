package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Ticket struct {
	Id              uuid.UUID
	TripID          uuid.UUID
	UserID          *uuid.UUID
	AgencyID        *uuid.UUID
	Count           uint
	Price           float64
	Status          string
	InvoiceId       uuid.UUID
	OwnerOfAgencyId uuid.UUID
	FactorId        string
}

func (t *Ticket) Validate() error {
	if t.Id == uuid.Nil {
		return errors.New("ticket ID can't be nil")
	}
	if t.TripID == uuid.Nil {
		return errors.New("trip ID can't be nil")
	}
	if t.UserID == nil {
		return errors.New("user ID can't be nil")
	}
	if t.AgencyID == nil {
		return errors.New("agency ID can't be nil")
	}
	if t.Count == 0 {
		return errors.New("count must be greater than zero")
	}
	if t.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if t.Status == "" {
		return errors.New("status can't be empty")
	}
	if t.InvoiceId == uuid.Nil {
		return errors.New("invoice ID can't be nil")
	}
	if t.OwnerOfAgencyId == uuid.Nil {
		return errors.New("owner of agency ID can't be nil")
	}
	return nil
}
