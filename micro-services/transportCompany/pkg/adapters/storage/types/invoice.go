package types

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	TicketID       uint      `gorm:"not null"`
	IssuedDate     time.Time `gorm:"not null"`
	Info           string    `gorm:"type:text"`
	PerAmountPrice float64
	TotalPrice     float64
	Status         string  `gorm:"type:varchar(20);default:'pending'"`
	Penalty        float64 `gorm:"default:0"`
}
