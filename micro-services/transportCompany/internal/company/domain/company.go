package domain

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type CompanyId uint

type Company struct {
	Id          uuid.UUID
	Name        string
	Description string
	OwnerId     uuid.UUID
	Address     string
	Phone       string
	Email       string
	CreatedAt   time.Time
	DeletedAt   time.Time
	UpdatedAt   time.Time
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("Name cant be empty")
	}
	if c.OwnerId == uuid.Nil {
		return errors.New("Owner id cant be nil")
	}
	if c.Address == "" {
		return errors.New("address cant be empty")
	}

	// Validate email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatched, err := regexp.MatchString(emailRegex, c.Email)
	if err != nil {
		return err
	}
	if !emailMatched {
		return errors.New("invalid email format")
	}

	// Validate Iranian landline number
	landlineRegex := `^0[1-8][0-9]{9}$`
	landlineMatched, err := regexp.MatchString(landlineRegex, c.Phone)
	if err != nil {
		return err
	}
	if !landlineMatched {
		return errors.New("invalid Iranian landline number format")
	}

	return nil
}

type CompanyFilter struct {
	Id   uuid.UUID
	Name string
}
