package staff

import (
	"context"
	"errors"
	"fmt"

	"gholi-fly-agancy/internal/staff/domain"
	"gholi-fly-agancy/internal/staff/port"

	"github.com/google/uuid"
)

var (
	ErrStaffCreateFailed     = errors.New("failed to create staff")
	ErrStaffNotFound         = errors.New("staff not found")
	ErrStaffValidationFailed = errors.New("staff validation failed")
)

type service struct {
	repo port.StaffRepo
}

func NewService(repo port.StaffRepo) port.StaffService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateStaff(ctx context.Context, staff domain.Staff) (domain.StaffID, error) {
	if err := staff.Validate(); err != nil {
		return domain.StaffID{}, fmt.Errorf("%w: %v", ErrStaffValidationFailed, err)
	}

	staffID, err := s.repo.Create(ctx, staff)
	if err != nil {
		return domain.StaffID{}, fmt.Errorf("%w: %v", ErrStaffCreateFailed, err)
	}

	return staffID, nil
}

func (s *service) GetStaffByID(ctx context.Context, id domain.StaffID) (*domain.Staff, error) {
	staff, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if staff == nil {
		return nil, ErrStaffNotFound
	}

	return staff, nil
}

func (s *service) UpdateStaff(ctx context.Context, staff domain.Staff) error {
	if err := staff.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrStaffValidationFailed, err)
	}

	if err := s.repo.Update(ctx, staff); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteStaff(ctx context.Context, id domain.StaffID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ListStaffByAgency(ctx context.Context, agencyID uuid.UUID) ([]domain.Staff, error) {
	staffList, err := s.repo.ListByAgencyID(ctx, agencyID)
	if err != nil {
		return nil, err
	}

	return staffList, nil
}
