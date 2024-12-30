package domain

import (
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
}
