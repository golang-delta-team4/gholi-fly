package storage

import (
	"context"
	"errors"

	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	"gholi-fly-hotel/internal/staff/domain"
	staffPort "gholi-fly-hotel/internal/staff/port"
	"gholi-fly-hotel/pkg/adapters/storage/mapper"
	"gholi-fly-hotel/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type staffRepo struct {
	db *gorm.DB
}

func NewStaffRepo(db *gorm.DB) staffPort.Repo {
	return &staffRepo{db: db}
}

func (r *staffRepo) CreateByHotelID(ctx context.Context, staffDomain domain.Staff, hotelID hotelDomain.HotelUUID) (domain.StaffUUID, error) {
	staff := mapper.StaffDomain2Storage(staffDomain)
	staff.HotelID = hotelID
	err := r.db.Table("staffs").WithContext(ctx).Create(staff).Error
	if err != nil {
		return domain.StaffUUID{}, err
	}
	return domain.StaffUUID(staff.UUID), nil
}

func (r *staffRepo) GetByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]domain.Staff, error) {
	var staffs []types.Staff
	err := r.db.Table("staffs").WithContext(ctx).Where("hotel_id = ?", hotelID).Find(&staffs).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchStaffStorage2Domain(staffs), nil
}

func (r *staffRepo) GetByID(ctx context.Context, staffID domain.StaffUUID) (*domain.Staff, error) {
	var staff types.Staff

	err := r.db.Table("staffs").WithContext(ctx).Where("uuid = ?", staffID).First(&staff).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if staff.ID == 0 {
		return nil, nil
	}

	return mapper.StaffStorage2Domain(staff), nil
}

func (r *staffRepo) Update(ctx context.Context, staffDomain domain.Staff) error {
	staff := mapper.StaffDomain2Storage(staffDomain)
	return r.db.Table("staffs").WithContext(ctx).Save(staff).Error
}

func (r *staffRepo) Delete(ctx context.Context, staffID domain.StaffUUID) error {
	return r.db.Table("staffs").WithContext(ctx).Delete(&types.Staff{}, "uuid = ?", staffID).Error
}
