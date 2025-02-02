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
	var trip types.Trip
	err := r.db.Table("trips").WithContext(ctx).Where("id = ?", id).First(&trip).Error
	if err != nil {
		return nil, err
	}

	domainTrip := mapper.TripStorage2Domain(trip)
	return domainTrip, nil
}

func (r *tripRepo) UpdateTrip(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	if err := r.db.WithContext(ctx).Debug().
		Model(&types.Trip{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update trip: %w", err)
	}

	return nil
}

func (r *tripRepo) GetTrips(ctx context.Context, pageSize int, page int, blockedUser []string) ([]domain.Trip, error) {
	var trips []types.Trip
	query := r.db.Table("trips").Joins("left join companies c on c.id = trips.company_id").WithContext(ctx).Limit(pageSize).Offset(page - 1*pageSize)
	if len(blockedUser) > 0 {
		query = query.Where("c.owner_id not in ?", blockedUser)
	}
	err := query.Find(&trips).Error
	if err != nil {
		return nil, err
	}
	var tripsDomain []domain.Trip
	for _, item := range trips {
		domainItem := mapper.TripStorage2Domain(item)
		tripsDomain = append(tripsDomain, *domainItem)
	}
	return tripsDomain, nil
}

func (r *tripRepo) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	return r.db.Table("trips").WithContext(ctx).Where("id = ?", id).Delete(&types.Trip{}).Error
}

func (r *tripRepo) ConfirmTrip(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Model(&types.Trip{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_confirmed": "true",
		}).Error; err != nil {
		return fmt.Errorf("failed to confirm trip: %w", err)
	}

	return nil
}
