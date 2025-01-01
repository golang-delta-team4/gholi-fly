package agency

import (
	"context"
	"errors"
	"fmt"

	"gholi-fly-agancy/internal/agency/domain"
	"gholi-fly-agancy/internal/agency/port"
)

var (
	ErrAgencyCreateFailed     = errors.New("failed to create agency")
	ErrAgencyNotFound         = errors.New("agency not found")
	ErrAgencyValidationFailed = errors.New("agency validation failed")
)

type service struct {
	repo port.AgencyRepo
}

func NewService(repo port.AgencyRepo) port.AgencyService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateAgency(ctx context.Context, agency domain.Agency) (domain.AgencyID, error) {
	if err := agency.Validate(); err != nil {
		return domain.AgencyID{}, fmt.Errorf("%w: %v", ErrAgencyValidationFailed, err)
	}

	agencyID, err := s.repo.Create(ctx, agency)
	if err != nil {
		return domain.AgencyID{}, fmt.Errorf("%w: %v", ErrAgencyCreateFailed, err)
	}

	return agencyID, nil
}

func (s *service) GetAgencyByID(ctx context.Context, id domain.AgencyID) (*domain.Agency, error) {
	agency, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if agency == nil {
		return nil, ErrAgencyNotFound
	}

	return agency, nil
}

func (s *service) UpdateAgency(ctx context.Context, agency domain.Agency) error {
	if err := agency.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrAgencyValidationFailed, err)
	}

	if err := s.repo.Update(ctx, agency); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteAgency(ctx context.Context, id domain.AgencyID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
