package factor

import (
	"context"
	"errors"
	"fmt"
	"gholi-fly-bank/internal/factor/domain"
	"gholi-fly-bank/internal/factor/port"
	"log"
)

var (
	ErrFactorCreation       = errors.New("error on creating factor")
	ErrFactorValidation     = errors.New("factor validation failed")
	ErrFactorNotFound       = errors.New("factor not found")
	ErrFactorStatusUpdate   = errors.New("error updating factor status")
	ErrInvalidFactorAmount  = errors.New("invalid factor amount")
	ErrInvalidSourceService = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

// NewService creates a new instance of the factor service.
func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateFactor(ctx context.Context, factor domain.Factor) (domain.FactorUUID, error) {
	// Validate factor details
	if factor.Amount <= 0 {
		return domain.FactorUUID{}, fmt.Errorf("%w: amount must be greater than zero", ErrInvalidFactorAmount)
	}
	if factor.SourceService == "" {
		return domain.FactorUUID{}, ErrInvalidSourceService
	}

	factorID, err := s.repo.Create(ctx, factor)
	if err != nil {
		log.Println("error creating factor:", err.Error())
		return domain.FactorUUID{}, ErrFactorCreation
	}

	return factorID, nil
}

func (s *service) GetFactorByID(ctx context.Context, factorID domain.FactorUUID) (*domain.Factor, error) {
	factor, err := s.repo.GetByID(ctx, factorID)
	if err != nil {
		log.Println("error fetching factor by ID:", err.Error())
		return nil, err
	}

	if factor == nil {
		return nil, ErrFactorNotFound
	}

	return factor, nil
}

func (s *service) GetFactors(ctx context.Context, filters domain.FactorFilters) ([]domain.Factor, int, error) {
	factors, total, err := s.repo.Get(ctx, filters)
	if err != nil {
		log.Println("error fetching factors:", err.Error())
		return nil, 0, err
	}

	return factors, total, nil
}

func (s *service) UpdateFactorStatus(ctx context.Context, factorID domain.FactorUUID, status domain.FactorStatus) error {
	// Validate the status transition if necessary
	if status == domain.FactorStatusUnknown {
		return fmt.Errorf("%w: invalid status update", ErrFactorStatusUpdate)
	}

	err := s.repo.UpdateStatus(ctx, factorID, status)
	if err != nil {
		log.Println("error updating factor status:", err.Error())
		return ErrFactorStatusUpdate
	}

	return nil
}
