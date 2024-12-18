package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

func CompanyDomain2Storage(companyDomain domain.Company) *types.Company {
	return &types.Company{
		Model: gorm.Model{
			ID: uint(companyDomain.Id),
		},
		Name:        companyDomain.Name,
		Description: companyDomain.Description,
		OwnerId:     companyDomain.OwnerId,
	}
}
