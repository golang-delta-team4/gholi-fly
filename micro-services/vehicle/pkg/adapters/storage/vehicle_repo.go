package storage

import (
	"context"

	"vehicle/internal/vehicle/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleRepo struct {
	db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) *VehicleRepo {
	return &VehicleRepo{db: db}
}

func (r *VehicleRepo) Create(ctx context.Context, vehicle *domain.Vehicle) error {
	return r.db.WithContext(ctx).Create(vehicle).Error
}

func (r *VehicleRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Vehicle, error) {
	var vehicle domain.Vehicle
	if err := r.db.WithContext(ctx).First(&vehicle, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *VehicleRepo) GetAll(ctx context.Context) ([]domain.Vehicle, error) {
	var vehicles []domain.Vehicle
	if err := r.db.WithContext(ctx).Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (r *VehicleRepo) Update(ctx context.Context, vehicle *domain.Vehicle) error {
	return r.db.WithContext(ctx).Save(vehicle).Error
}

func (r *VehicleRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Vehicle{}, "id = ?", id).Error
}

func (r *VehicleRepo) ProcessTripRequest(ctx context.Context) (*domain.TripRequest, error) {

	// implementation
	return nil, nil
}
