package toure

import (
	"context"
	"errors"
	"fmt"

	"gholi-fly-agancy/internal/toure/domain"
	"gholi-fly-agancy/internal/toure/port"

	"github.com/google/uuid"
)

var (
	ErrToureCreateFailed     = errors.New("failed to create toure")
	ErrToureNotFound         = errors.New("toure not found")
	ErrToureValidationFailed = errors.New("toure validation failed")
)

type service struct {
	repo port.ToureRepo
}

func NewService(repo port.ToureRepo) port.ToureService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateToure(ctx context.Context, toure domain.Toure) (domain.ToureID, error) {
	if err := toure.Validate(); err != nil {
		return domain.ToureID{}, fmt.Errorf("%w: %v", ErrToureValidationFailed, err)
	}

	toureID, err := s.repo.Create(ctx, toure)
	if err != nil {
		return domain.ToureID{}, fmt.Errorf("%w: %v", ErrToureCreateFailed, err)
	}

	return toureID, nil
}

func (s *service) GetToureByID(ctx context.Context, id domain.ToureID) (*domain.Toure, error) {
	toure, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if toure == nil {
		return nil, ErrToureNotFound
	}

	return toure, nil
}

func (s *service) UpdateToure(ctx context.Context, toure domain.Toure) error {
	if err := toure.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrToureValidationFailed, err)
	}

	if err := s.repo.Update(ctx, toure); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteToure(ctx context.Context, id domain.ToureID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ListTouresByAgency(ctx context.Context, agencyID uuid.UUID) ([]domain.Toure, error) {
	toures, err := s.repo.ListByAgencyID(ctx, agencyID)
	if err != nil {
		return nil, err
	}

	return toures, nil
}
