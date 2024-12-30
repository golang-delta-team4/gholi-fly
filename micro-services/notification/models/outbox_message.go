package models

import (
	uuid "github.com/google/uuid"
)

type OutBoxMessage struct {
	ID          string    `gorm:"id" json:"id"`
	EventName   string    `gorm:"event_name" json:"event_name"`
	UserID      uuid.UUID `gorm:"user_id" json:"user_id"`
	Name        string    `gorm:"name" json:"name"`
	Email       string    `gorm:"email" json:"email"`
	Message     string    `gorm:"message" json:"message"`
	IsProcessed bool      `gorm:"is_processed" json:"is_processed"`
}
