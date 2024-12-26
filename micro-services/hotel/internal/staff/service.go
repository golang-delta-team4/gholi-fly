package staff

import (
	"context"
	"errors"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	staffDomain "gholi-fly-hotel/internal/staff/domain"
	"gholi-fly-hotel/internal/staff/port"
	"strings"
)

var (
	ErrStaffCreation           = errors.New("error on creating staff")
	ErrStaffCreationValidation = errors.New("error on creating staff: validation failed")
	ErrStaffCreationDuplicate  = errors.New("staff already exists")
	ErrStaffNotFound           = errors.New("staff not found")
	ErrInvalidSourceService    = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// CreateStaff creates a new staff
func (s *service) CreateStaffByHotelID(ctx context.Context, staff staffDomain.Staff, hotelID hotelDomain.HotelUUID) (staffDomain.StaffUUID, error) {
	staffID, err := s.repo.CreateByHotelID(ctx, staff, hotelID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return staffDomain.StaffUUID{}, ErrStaffCreationDuplicate
		}
		return staffDomain.StaffUUID{}, ErrStaffCreation
	}
	return staffID, nil
}

// GetStaffByID returns a staff by its ID
func (s *service) GetStaffByID(ctx context.Context, staffID staffDomain.StaffUUID) (*staffDomain.Staff, error) {
	staff, err := s.repo.GetByID(ctx, staffID)
	if err != nil {
		return nil, ErrStaffNotFound
	}
	return staff, nil
}

// GetStaffs returns all staffs
func (s *service) GetAllStaffsByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]staffDomain.Staff, error) {
	staffs, err := s.repo.GetByHotelID(ctx, hotelID)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}

// UpdateStaff updates a staff
func (s *service) UpdateStaff(ctx context.Context, staff staffDomain.Staff) error {
	err := s.repo.Update(ctx, staff)
	if err != nil {
		return err
	}
	return nil
}

// DeleteStaff deletes a staff
func (s *service) DeleteStaff(ctx context.Context, staffID staffDomain.StaffUUID) error {
	err := s.repo.Delete(ctx, staffID)
	if err != nil {
		return err
	}
	return nil
}
