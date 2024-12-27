package storage

import (
	"context"
	"errors"

	"gholi-fly-agancy/internal/factor/domain"
	"gholi-fly-agancy/internal/factor/port"
	"gholi-fly-agancy/pkg/adapters/storage/mapper"
	"gholi-fly-agancy/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type factorRepo struct {
	db *gorm.DB
}

func NewFactorRepo(db *gorm.DB) port.FactorRepo {
	return &factorRepo{db}
}

func (r *factorRepo) Create(ctx context.Context, factorDomain domain.Factor) (domain.FactorID, error) {
	factor := mapper.FactorDomain2Storage(factorDomain)
	return domain.FactorID(factor.ID), r.db.Table("factors").WithContext(ctx).Create(factor).Error
}

func (r *factorRepo) GetByID(ctx context.Context, id domain.FactorID) (*domain.Factor, error) {
	var factor types.Factor
	err := r.db.Table("factors").WithContext(ctx).First(&factor, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.FactorStorage2Domain(factor), nil
}

func (r *factorRepo) Update(ctx context.Context, factorDomain domain.Factor) error {
	factor := mapper.FactorDomain2Storage(factorDomain)
	return r.db.Table("factors").WithContext(ctx).Save(factor).Error
}

func (r *factorRepo) Delete(ctx context.Context, id domain.FactorID) error {
	return r.db.Table("factors").WithContext(ctx).Delete(&types.Factor{}, "id = ?", id).Error
}

func (r *factorRepo) ListByReservationID(ctx context.Context, reservationID uuid.UUID) ([]domain.Factor, error) {
	var factors []types.Factor
	err := r.db.Table("factors").WithContext(ctx).Where("reservation_id = ?", reservationID).Find(&factors).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchFactorStorage2Domain(factors), nil
}
