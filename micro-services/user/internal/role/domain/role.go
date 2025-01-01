package domain

import (
	"time"
	"user-service/internal/permission/domain"

	"github.com/google/uuid"
)

type Role struct {
	UUID uuid.UUID
	Name string
	CreatedAt time.Time
	Permissions []domain.Permission
}