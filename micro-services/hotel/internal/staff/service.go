package staff

import (
	"context"
	"errors"
	"gholi-fly-hotel/internal/staff/domain"
	"gholi-fly-hotel/internal/staff/port"
)

var (
	ErrStaffCreation        = errors.New("error on creating staff")
	ErrStaffNotFound        = errors.New("staff not found")
	ErrInvalidSourceService = errors.New("invalid source service")
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
func (s *service) CreateStaff(ctx context.Context, staff domain.Staff) (domain.StaffUUID, error) {
	staffID, err := s.repo.Create(ctx, staff)
	if err != nil {
		return domain.StaffUUID{}, ErrStaffCreation
	}
	return staffID, nil
}

// GetStaffByID returns a staff by its ID
func (s *service) GetStaffByID(ctx context.Context, staffID domain.StaffUUID) (*domain.Staff, error) {
	staff, err := s.repo.GetByID(ctx, staffID)
	if err != nil {
		return nil, ErrStaffNotFound
	}
	return staff, nil
}

// GetStaffs returns all staffs
func (s *service) GetStaffs(ctx context.Context) ([]domain.Staff, error) {
	staffs, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}

// UpdateStaff updates a staff
func (s *service) UpdateStaff(ctx context.Context, staff domain.Staff) error {
	err := s.repo.Update(ctx, staff)
	if err != nil {
		return err
	}
	return nil
}

// DeleteStaff deletes a staff
func (s *service) DeleteStaff(ctx context.Context, staffID domain.StaffUUID) error {
	err := s.repo.Delete(ctx, staffID)
	if err != nil {
		return err
	}
	return nil
}
