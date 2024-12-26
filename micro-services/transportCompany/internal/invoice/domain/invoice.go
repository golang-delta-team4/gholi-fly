package domain

import (
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
