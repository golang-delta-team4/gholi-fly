package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CompanyId uint

type Company struct {
	Id          CompanyId
	Name        string
	Description string
	OwnerId     uuid.UUID
	CreatedAt   time.Time
	DeletedAt   time.Time
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("Name cant be empty")
	}
	if c.OwnerId == uuid.Nil {
		return errors.New("Owner id cant be nil")
	}

	return nil
}

type CompanyFilter struct {
	Id   CompanyId
	Name string
}
