package storage

import (
	"context"
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/mapper"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tripRepo struct {
	db *gorm.DB
}

func NewTripRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	repo := &tripRepo{db}
	return repo
}

func (r *tripRepo) CreateTrip(ctx context.Context, tripDomain domain.Trip) (uuid.UUID, error) {
	trip := mapper.TripDomain2Storage(tripDomain)
	return trip.Id, r.db.Table("trips").WithContext(ctx).Create(trip).Error
}

func (r *tripRepo) GetTripById(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	err := r.db.Table("trips").WithContext(ctx).Where("id = ?", id).First(&trip).Error
	if err != nil {
		return nil, err
	}
	return &trip, nil
}

func (r *tripRepo) UpdateTrip(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	if err := r.db.WithContext(ctx).
		Model(&types.Trip{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update trip: %w", err)
	}

	return nil
}

func (r *tripRepo) GetTrips(ctx context.Context, pageSize int, page int) ([]domain.Trip, error) {
	var trips []domain.Trip
	err := r.db.Table("trips").WithContext(ctx).Limit(pageSize).Offset(page - 1*pageSize).Find(&trips).Error
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *tripRepo) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	return r.db.Table("trips").WithContext(ctx).Where("id = ?", id).Delete(&types.Trip{}).Error
}
