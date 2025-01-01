package storage

import (
	"context"
	"errors"

	"gholi-fly-agancy/internal/staff/domain"
	"gholi-fly-agancy/internal/staff/port"
	"gholi-fly-agancy/pkg/adapters/storage/mapper"
	"gholi-fly-agancy/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type staffRepo struct {
	db *gorm.DB
}

func NewStaffRepo(db *gorm.DB) port.StaffRepo {
	return &staffRepo{db}
}

func (r *staffRepo) Create(ctx context.Context, staffDomain domain.Staff) (domain.StaffID, error) {
	staff := mapper.StaffDomain2Storage(staffDomain)
	return domain.StaffID(staff.ID), r.db.Table("staffs").WithContext(ctx).Create(staff).Error
}

func (r *staffRepo) GetByID(ctx context.Context, id domain.StaffID) (*domain.Staff, error) {
	var staff types.Staff
	err := r.db.Table("staffs").WithContext(ctx).First(&staff, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.StaffStorage2Domain(staff), nil
}

func (r *staffRepo) Update(ctx context.Context, staffDomain domain.Staff) error {
	staff := mapper.StaffDomain2Storage(staffDomain)
	return r.db.Table("staffs").WithContext(ctx).Save(staff).Error
}

func (r *staffRepo) Delete(ctx context.Context, id domain.StaffID) error {
	return r.db.Table("staffs").WithContext(ctx).Delete(&types.Staff{}, "id = ?", id).Error
}

func (r *staffRepo) ListByAgencyID(ctx context.Context, agencyID uuid.UUID) ([]domain.Staff, error) {
	var staffs []types.Staff
	err := r.db.Table("staffs").WithContext(ctx).Where("agency_id = ?", agencyID).Find(&staffs).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchStaffStorage2Domain(staffs), nil
}
