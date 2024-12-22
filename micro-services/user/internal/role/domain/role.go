package domain

import (
	"user-service/internal/permission/domain"

	"github.com/google/uuid"
)

type Role struct {
	UUID uuid.UUID
	Name string
	Permissions []domain.Permission
}