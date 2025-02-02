package storage

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/mapper"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	repo := &companyRepo{db}

	if !cached {
		return repo
	}

	// return &companyCashedRepo{
	// 	repo:     repo,
	// 	provider: provider,
	// }

	return repo
}

func (r *companyRepo) Create(ctx context.Context, companyDomain domain.Company) (uuid.UUID, error) {
	company := mapper.CompanyDomain2Storage(companyDomain)
	return company.Id, r.db.Table("companies").WithContext(ctx).Create(company).Error
}

func (r *companyRepo) GetCompanyById(ctx context.Context, companyId uuid.UUID) (*domain.Company, error) {
	var company types.Company
	result := r.db.Table("companies").WithContext(ctx).First(&company, "id = ?", companyId)
	if result.Error != nil {
		return nil, result.Error
	}
	companyDomain := mapper.CompanyStorage2Domain(company)

	return companyDomain, nil
}

func (r *companyRepo) GetByOwnerId(ctx context.Context, ownerId uuid.UUID) (*domain.Company, error) {
	var company types.Company
	result := r.db.Table("companies").WithContext(ctx).First(&company, "owner_id = ?", ownerId)
	if result.Error != nil {
		return nil, result.Error
	}
	companyDomain := mapper.CompanyStorage2Domain(company)

	return companyDomain, nil
}

func (r *companyRepo) UpdateCompany(ctx context.Context, company domain.Company) error {
	updates := make(map[string]interface{})

	if company.Name != "" {
		updates["name"] = company.Name
	}
	if company.Description != "" {
		updates["description"] = company.Description
	}
	if company.Address != "" {
		updates["address"] = company.Address
	}
	if company.Phone != "" {
		updates["phone"] = company.Phone
	}
	if company.Email != "" {
		updates["email"] = company.Email
	}

	return r.db.Model(&types.Company{}).Where("id = ?", company.Id).Updates(updates).Error
}

func (r *companyRepo) DeleteCompany(ctx context.Context, companyId uuid.UUID) error {
	return r.db.Table("companies").WithContext(ctx).Delete(&types.Company{}, "id = ?", companyId).Error
}
