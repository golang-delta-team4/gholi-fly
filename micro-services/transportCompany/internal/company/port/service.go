package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/google/uuid"
)

type Service interface {
	CreateCompany(ctx context.Context, company domain.Company) (uuid.UUID, error)
	GetCompanyById(ctx context.Context, companyId uuid.UUID) (*domain.Company, error)
	GetByOwnerId(ctx context.Context, ownerId uuid.UUID) (*domain.Company, error)
}
