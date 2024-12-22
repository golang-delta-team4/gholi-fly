package trip

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/google/uuid"
)

var (
	ErrTripOnCreate           = errors.New("error on creating new trip")
	ErrTripCreationValidation = errors.New("validation failed")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, trip domain.Trip) (uuid.UUID, error) {
	if err := trip.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}
	companyId, err := s.repo.CreateTrip(ctx, trip)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
		return uuid.Nil, err
	}

	return companyId, nil
}
