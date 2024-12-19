package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	Id         uuid.UUID `gorm:"type:uuid;;primaryKey"`
	IssuedDate time.Time `gorm:"not null"`
	Info       string    `gorm:"type:text"`
	TotalPrice float64
	Status     string `gorm:"default:'pending'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
