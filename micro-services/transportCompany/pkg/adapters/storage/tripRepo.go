package storage

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/mapper"
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
