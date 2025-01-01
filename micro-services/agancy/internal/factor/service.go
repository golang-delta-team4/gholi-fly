package factor

import (
	"context"
	"errors"
	"fmt"

	"gholi-fly-agancy/internal/factor/domain"
	"gholi-fly-agancy/internal/factor/port"

	"github.com/google/uuid"
)

var (
	ErrFactorCreateFailed     = errors.New("failed to create factor")
	ErrFactorNotFound         = errors.New("factor not found")
	ErrFactorValidationFailed = errors.New("factor validation failed")
)

type service struct {
	repo port.FactorRepo
}

func NewService(repo port.FactorRepo) port.FactorService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateFactor(ctx context.Context, factor domain.Factor) (domain.FactorID, error) {
	if err := factor.Validate(); err != nil {
		return domain.FactorID{}, fmt.Errorf("%w: %v", ErrFactorValidationFailed, err)
	}

	factorID, err := s.repo.Create(ctx, factor)
	if err != nil {
		return domain.FactorID{}, fmt.Errorf("%w: %v", ErrFactorCreateFailed, err)
	}

	return factorID, nil
}

func (s *service) GetFactorByID(ctx context.Context, id domain.FactorID) (*domain.Factor, error) {
	factor, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if factor == nil {
		return nil, ErrFactorNotFound
	}

	return factor, nil
}

func (s *service) UpdateFactor(ctx context.Context, factor domain.Factor) error {
	if err := factor.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrFactorValidationFailed, err)
	}

	if err := s.repo.Update(ctx, factor); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteFactor(ctx context.Context, id domain.FactorID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ListFactorsByReservation(ctx context.Context, reservationID uuid.UUID) ([]domain.Factor, error) {
	factors, err := s.repo.ListByReservationID(ctx, reservationID)
	if err != nil {
		return nil, err
	}

	return factors, nil
}
