package port

import (
	"context"

	"gholi-fly-agancy/internal/toure/domain"

	"github.com/google/uuid"
)

type ToureService interface {
	CreateToure(ctx context.Context, toure domain.Toure) (domain.ToureID, error)
	GetToureByID(ctx context.Context, id domain.ToureID) (*domain.Toure, error)
	UpdateToure(ctx context.Context, toure domain.Toure) error
	DeleteToure(ctx context.Context, id domain.ToureID) error
	ListTouresByAgency(ctx context.Context, agencyID uuid.UUID) ([]domain.Toure, error)
}
