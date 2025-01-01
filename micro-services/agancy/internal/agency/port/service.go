package port

import (
	"context"

	"gholi-fly-agancy/internal/agency/domain"
)

type AgencyService interface {
	CreateAgency(ctx context.Context, agency domain.Agency) (domain.AgencyID, error)
	GetAgencyByID(ctx context.Context, id domain.AgencyID) (*domain.Agency, error)
	UpdateAgency(ctx context.Context, agency domain.Agency) error
	DeleteAgency(ctx context.Context, id domain.AgencyID) error
}
