package tour

import (
	"context"
	"errors"
	"fmt"
	"gholi-fly-agancy/internal/tour/domain"
	"gholi-fly-agancy/internal/tour/port"

	"github.com/google/uuid"
)

var (
	ErrtourCreateFailed     = errors.New("failed to create tour")
	ErrtourNotFound         = errors.New("tour not found")
	ErrtourValidationFailed = errors.New("tour validation failed")
)

type service struct {
	repo port.TourRepo
}

func NewService(repo port.TourRepo) port.TourService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error) {
	if err := tour.Validate(); err != nil {
		return domain.TourID{}, fmt.Errorf("%w: %v", ErrtourValidationFailed, err)
	}

	tourID, err := s.repo.Create(ctx, tour)
	if err != nil {
		return domain.TourID{}, fmt.Errorf("%w: %v", ErrtourCreateFailed, err)
	}

	return tourID, nil
}

func (s *service) GetTourByID(ctx context.Context, id domain.TourID) (*domain.Tour, error) {
	tour, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tour == nil {
		return nil, ErrtourNotFound
	}

	return tour, nil
}

func (s *service) UpdateTour(ctx context.Context, tour domain.Tour) error {
	if err := tour.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrtourValidationFailed, err)
	}

	if err := s.repo.Update(ctx, tour); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteTour(ctx context.Context, id domain.TourID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ListToursByAgency(ctx context.Context, agencyID uuid.UUID) ([]domain.Tour, error) {
	tours, err := s.repo.ListByAgencyID(ctx, agencyID)
	if err != nil {
		return nil, err
	}

	return tours, nil
}
