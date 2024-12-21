package storage

import (
	"context"
	"log"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/google/uuid"
)

type companyCashedRepo struct {
	repo     companyPort.Repo
	provider cache.Provider
}

func (r *companyCashedRepo) Create(ctx context.Context, companyDomain domain.Company) (uuid.UUID, error) {
	uId, err := r.repo.Create(ctx, companyDomain)
	if err != nil {
		return uuid.Nil, err
	}
	companyDomain.Id = uId

	oc := cache.NewJsonObjectCacher[*domain.Company](r.provider)
	if err := oc.Set(ctx, r.companyFilterKey(&domain.CompanyFilter{
		Id: uId,
	}), 0, &companyDomain); err != nil {
		log.Println("error on caching (SET) company with id :", uId)
	}

	return uId, nil
}

func (r *companyCashedRepo) GetCompanyById(ctx context.Context, companyId uuid.UUID) (*domain.Company, error) {
	return nil, nil
}

func (r *companyCashedRepo) companyFilterKey(filter *domain.CompanyFilter) string {
	return "companies." + filter.Id.String() + "." + filter.Name
}
