package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, company domain.Company) (uuid.UUID, error)
	GetCompanyById(ctx context.Context, companyId uuid.UUID) (*domain.Company, error)
}
