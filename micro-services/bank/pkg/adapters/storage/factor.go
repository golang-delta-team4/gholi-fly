package storage

import (
	"context"
	"errors"
	"gholi-fly-bank/internal/factor/domain"
	"gholi-fly-bank/internal/factor/port"
	"gholi-fly-bank/pkg/adapters/storage/mapper"
	"gholi-fly-bank/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type factorRepo struct {
	db *gorm.DB
}

// NewFactorRepo creates a new instance of the factor repository.
func NewFactorRepo(db *gorm.DB) port.Repo {
	return &factorRepo{db: db}
}

func (r *factorRepo) Create(ctx context.Context, factorDomain domain.Factor) (domain.FactorUUID, error) {
	factor := mapper.FactorDomain2Storage(factorDomain)
	err := r.db.WithContext(ctx).Table("factors").Create(factor).Error
	if err != nil {
		return domain.FactorUUID{}, err
	}
	return domain.FactorUUID(factor.ID), nil
}

func (r *factorRepo) GetByID(ctx context.Context, factorID domain.FactorUUID) (*domain.Factor, error) {
	var factor types.Factor
	err := r.db.WithContext(ctx).Table("factors").Where("id = ?", factorID).First(&factor).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.FactorStorage2Domain(factor), nil
}

func (r *factorRepo) Get(ctx context.Context, filters domain.FactorFilters) ([]domain.Factor, error) {
	var factors []types.Factor
	query := r.db.WithContext(ctx).Table("factors")

	// Apply filters
	if filters.SourceService != "" {
		query = query.Where("source_service = ?", filters.SourceService)
	}
	if filters.BookingID != "" {
		query = query.Where("booking_id = ?", filters.BookingID)
	}
	if filters.CustomerID != uuid.Nil {
		query = query.Where("customer_id = ?", filters.CustomerID)
	}
	if filters.Status > 0 {
		query = query.Where("status = ?", uint8(filters.Status))
	}

	err := query.Find(&factors).Error
	if err != nil {
		return nil, err
	}

	return mapper.BatchFactorStorage2Domain(factors), nil
}

func (r *factorRepo) UpdateStatus(ctx context.Context, factorID domain.FactorUUID, status domain.FactorStatus) error {
	result := r.db.WithContext(ctx).Table("factors").
		Where("id = ?", factorID).
		Update("status", uint8(status))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("factor not found")
	}

	return nil
}
