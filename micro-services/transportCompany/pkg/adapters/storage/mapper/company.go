package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
)

func CompanyDomain2Storage(companyDomain domain.Company) *types.Company {
	return &types.Company{
		Id:          companyDomain.Id,
		Name:        companyDomain.Name,
		Description: companyDomain.Description,
		OwnerId:     companyDomain.OwnerId,
		Address:     companyDomain.Address,
		Phone:       companyDomain.Phone,
		Email:       companyDomain.Email,
	}
}
