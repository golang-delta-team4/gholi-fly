package storage

import (
	"context"

	"vehicle/internal/vehicle/domain"
	"vehicle/pkg/adapters/storage/types"

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

func (r *VehicleRepo) GetMatchedVehicle(ctx context.Context, vehicleMatchRequest *domain.MatchMakerRequest) (types.Vehicle, error) {
	var vehicle types.Vehicle
	subQuery := r.db.Model(&types.VehicleReserve{}).
		Table("vehicle_reserves as vr2").
		Select("count(*)").
		Where("vr2.vehicle_id = vehicles.id and (? >= vr2.start_date and ? < vr2.end_date) or (? < vr2.end_date and ? > vr2.start_date)",
			vehicleMatchRequest.ReserveStartDate,
			vehicleMatchRequest.ReserveStartDate,
			vehicleMatchRequest.ReserveEndDate,
			vehicleMatchRequest.ReserveEndDate)
	err := r.db.Model(&types.Vehicle{}).
		Joins("left join vehicle_reserves vr on vr.vehicle_id = vehicles.id").
		Select("vehicles.id, vehicles.owner_id, vehicles.type, vehicles.capacity, vehicles.speed, vehicles.unique_code, vehicles.status, vehicles.year_of_manufacture, vehicles.created_at, vehicles.updated_at, vehicles.price_per_kilometer, ? * price_per_kilometer as computed_price", vehicleMatchRequest.TripDistance).
		Where("0 = (?)", subQuery).
		Where("vehicles.type = ? AND vehicles.capacity >= ? AND ? * vehicles.price_per_kilometer < ? AND vehicles.year_of_manufacture = ?",
			vehicleMatchRequest.TripType,
			vehicleMatchRequest.NumberOfPassengers,
			vehicleMatchRequest.TripDistance,
			vehicleMatchRequest.MaxPrice,
			vehicleMatchRequest.YearOfManufacture).
		Order("vehicles.capacity asc, computed_price asc, vehicles.year_of_manufacture asc, vehicles.created_at desc").
		First(&vehicle).Error
	return vehicle, err
}

func (r *VehicleRepo) CreateReservation(ctx context.Context, vehicleReserve types.VehicleReserve) (uuid.UUID, error) {
	return vehicleReserve.ID, r.db.Model(&types.VehicleReserve{}).Create(&vehicleReserve).Error

}
