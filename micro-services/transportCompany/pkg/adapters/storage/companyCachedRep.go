package storage

import (
	"context"
	"log"
	"strconv"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
)

type companyCashedRepo struct {
	repo     companyPort.Repo
	provider cache.Provider
}

func (r *companyCashedRepo) Create(ctx context.Context, companyDomain domain.Company) (domain.CompanyId, error) {
	uId, err := r.repo.Create(ctx, companyDomain)
	if err != nil {
		return 0, err
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

func (r *companyCashedRepo) companyFilterKey(filter *domain.CompanyFilter) string {
	return "companies." + strconv.FormatUint(uint64(filter.Id), 10) + "." + filter.Name
}
