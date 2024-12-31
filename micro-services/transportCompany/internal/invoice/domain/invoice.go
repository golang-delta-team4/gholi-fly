package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	Id         uuid.UUID
	IssuedDate time.Time
	Info       string
	TotalPrice float64
	Status     uint8
}

const (
	Pending uint8 = iota
	Canceled
	Paid
)

func (i *Invoice) Validate() error {
	if i.Id == uuid.Nil {
		return errors.New("invoice ID can't be nil")
	}
	if i.IssuedDate.IsZero() {
		return errors.New("issued date can't be zero")
	}
	if i.IssuedDate.After(time.Now()) {
		return errors.New("issued date can't be in the future")
	}
	if i.Info == "" {
		return errors.New("info can't be empty")
	}
	if i.TotalPrice <= 0 {
		return errors.New("total price must be greater than zero")
	}
	if !isValidInvoiceStatus(i.Status) {
		return errors.New("status must be one of 'Pending', 'Canceled', or 'Paid'")
	}
	return nil
}

func isValidInvoiceStatus(status uint8) bool {
	switch status {
	case Pending, Canceled, Paid:
		return true
	default:
		return false
	}
}
