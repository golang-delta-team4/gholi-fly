package storage

import (
	"context"
	"errors"

	"gholi-fly-agancy/internal/tour/domain"
	"gholi-fly-agancy/internal/tour/port"
	"gholi-fly-agancy/pkg/adapters/storage/mapper"
	"gholi-fly-agancy/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tourRepo struct {
	db *gorm.DB
}

func NewTourRepo(db *gorm.DB) port.TourRepo {
	return &tourRepo{db}
}

func (r *tourRepo) Create(ctx context.Context, tourDomain domain.Tour) (domain.TourID, error) {
	tour := mapper.TourDomain2Storage(tourDomain)
	return domain.TourID(tour.ID), r.db.Table("tours").WithContext(ctx).Create(tour).Error
}

func (r *tourRepo) GetByID(ctx context.Context, id domain.TourID) (*domain.Tour, error) {
	var tour types.Tour
	err := r.db.Table("tours").WithContext(ctx).First(&tour, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.TourStorage2Domain(tour), nil
}

func (r *tourRepo) Update(ctx context.Context, tourDomain domain.Tour) error {
	tour := mapper.TourDomain2Storage(tourDomain)
	return r.db.Table("tours").WithContext(ctx).Save(tour).Error
}

func (r *tourRepo) Delete(ctx context.Context, id domain.TourID) error {
	return r.db.Table("tours").WithContext(ctx).Delete(&types.Tour{}, "id = ?", id).Error
}

func (r *tourRepo) ListByAgencyID(ctx context.Context, agencyID uuid.UUID) ([]domain.Tour, error) {
	var tours []types.Tour
	err := r.db.Table("tours").WithContext(ctx).Where("agency_id = ?", agencyID).Find(&tours).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchTourStorage2Domain(tours), nil
}
