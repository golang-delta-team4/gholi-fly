package storage

import (
	"context"
	"errors"

	"gholi-fly-agancy/internal/agency/domain"
	"gholi-fly-agancy/internal/agency/port"
	"gholi-fly-agancy/pkg/adapters/storage/mapper"
	"gholi-fly-agancy/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type agencyRepo struct {
	db *gorm.DB
}

func NewAgencyRepo(db *gorm.DB) port.AgencyRepo {
	return &agencyRepo{db}
}

func (r *agencyRepo) Create(ctx context.Context, agencyDomain domain.Agency) (domain.AgencyID, error) {
	agency := mapper.AgencyDomain2Storage(agencyDomain)
	return domain.AgencyID(agency.ID), r.db.Table("agencies").WithContext(ctx).Create(agency).Error
}

func (r *agencyRepo) GetByID(ctx context.Context, id domain.AgencyID) (*domain.Agency, error) {
	var agency types.Agency
	err := r.db.Table("agencies").WithContext(ctx).First(&agency, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.AgencyStorage2Domain(agency), nil
}

func (r *agencyRepo) Update(ctx context.Context, agencyDomain domain.Agency) error {
	agency := mapper.AgencyDomain2Storage(agencyDomain)
	return r.db.Table("agencies").WithContext(ctx).Save(agency).Error
}

func (r *agencyRepo) Delete(ctx context.Context, id domain.AgencyID) error {
	return r.db.Table("agencies").WithContext(ctx).Delete(&types.Agency{}, "id = ?", id).Error
}
