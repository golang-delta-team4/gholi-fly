package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type NotificationHistory struct {
	ID        string    `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"user_id"`
	Name      string    `gorm:"name"`
	Email     string    `gorm:"email"`
	Event     string    `gorm:"event"`
	Message   string    `gorm:"message"`
	IsRead    bool      `gorm:"is_read"`
	CreatedAt time.Time
}
