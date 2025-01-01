package port

import (
	"context"
	"gholi-fly-agancy/internal/tour/domain"

	"github.com/google/uuid"
)

type TourService interface {
	CreateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error)
	GetTourByID(ctx context.Context, id domain.TourID) (*domain.Tour, error)
	UpdateTour(ctx context.Context, tour domain.Tour) error
	DeleteTour(ctx context.Context, id domain.TourID) error
	ListToursByAgency(ctx context.Context, agencyID uuid.UUID) ([]domain.Tour, error)
}
