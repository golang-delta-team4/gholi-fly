package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
)

type Repo interface {
	Create(ctx context.Context, company domain.Company) (domain.CompanyId, error)
}
