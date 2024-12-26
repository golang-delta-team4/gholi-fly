package port

import (
	"context"
	"gholi-fly-agancy/internal/agency/domain"
)

type AgencyRepo interface {
	Create(ctx context.Context, agency domain.Agency) (domain.AgencyID, error)
	GetByID(ctx context.Context, id domain.AgencyID) (*domain.Agency, error)
	Update(ctx context.Context, agency domain.Agency) error
	Delete(ctx context.Context, id domain.AgencyID) error
}
