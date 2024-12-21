package domain

import "github.com/google/uuid"

type Permission struct {
	UUID       uuid.UUID
	Route      string
	Method     string
	ResourceID *uint
}
