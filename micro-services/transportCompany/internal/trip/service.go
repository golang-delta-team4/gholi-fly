package trip

import (
	"context"
	"errors"
	"fmt"
	"log"

	technicalTeamPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	tripRepo "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/google/uuid"
)

var (
	ErrTripOnCreate           = errors.New("error on creating new trip")
	ErrCanNotUpdate           = errors.New("can not update")
	ErrTripCreationValidation = errors.New("validation failed")
)

type service struct {
	repo              port.Repo
	technicalTeamRepo technicalTeamPort.Repo
	tripRepo          tripRepo.Repo
}

func NewService(repo port.Repo, technicalTeamRepo technicalTeamPort.Repo, tripRepo tripRepo.Repo) port.Service {
	return &service{
		repo:              repo,
		technicalTeamRepo: technicalTeamRepo,
		tripRepo:          tripRepo,
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

func (s *service) GetTripById(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	trip, err := s.repo.GetTripById(ctx, id)
	if err != nil {
		log.Println("error on getting trip by id: ", err.Error())
		return nil, err
	}

	return trip, nil
}

func (s *service) UpdateTrip(ctx context.Context, newTrip domain.Trip, oldTrip domain.Trip) error {
	updates := make(map[string]interface{})

	if newTrip.UserPrice != 0 && newTrip.AgencyPrice == 0 {
		if newTrip.UserPrice > oldTrip.AgencyPrice {
			updates["user_price"] = newTrip.UserPrice
		} else {
			return ErrCanNotUpdate
		}
	}

	if newTrip.UserPrice == 0 && newTrip.AgencyPrice != 0 {
		if oldTrip.UserPrice > newTrip.AgencyPrice {
			updates["agency_price"] = newTrip.UserPrice
		} else {
			return ErrCanNotUpdate
		}
	}

	if newTrip.UserPrice != 0 && newTrip.AgencyPrice != 0 {
		if newTrip.UserPrice > newTrip.AgencyPrice {
			updates["agency_price"] = newTrip.AgencyPrice
			updates["user_price"] = newTrip.UserPrice
		} else {
			return ErrCanNotUpdate
		}
	}

	if newTrip.UserReleaseDate.IsZero() && !newTrip.TourReleaseDate.IsZero() {
		if oldTrip.UserReleaseDate.After(newTrip.TourReleaseDate) {
			updates["tour_release_date"] = newTrip.TourReleaseDate
		} else {
			return ErrCanNotUpdate
		}
	}

	if !newTrip.UserReleaseDate.IsZero() && newTrip.TourReleaseDate.IsZero() {
		if newTrip.UserReleaseDate.After(oldTrip.TourReleaseDate) {
			updates["user_release_date"] = newTrip.UserReleaseDate
		} else {
			return ErrCanNotUpdate
		}
	}

	if !newTrip.UserReleaseDate.IsZero() && !newTrip.TourReleaseDate.IsZero() {
		if newTrip.UserReleaseDate.After(newTrip.TourReleaseDate) {
			updates["user_release_date"] = newTrip.UserReleaseDate
			updates["tour_release_date"] = newTrip.TourReleaseDate
		} else {
			return ErrCanNotUpdate
		}
	}

	if newTrip.MinPassengers != 0 {
		if oldTrip.VehicleID != nil {
			return ErrCanNotUpdate
		} else {
			updates["min_passengers"] = newTrip.MinPassengers
		}
	}

	if newTrip.TechnicalTeamID != nil {
		updates["tech_team_id"] = newTrip.TechnicalTeamID
	}

	if newTrip.MaxTickets != 0 {
		if oldTrip.VehicleID != nil {
			return ErrCanNotUpdate
		} else {
			updates["max_tickets"] = newTrip.MaxTickets
		}
	}

	if newTrip.IsCanceled {
		return ErrCanNotUpdate
	}

	if newTrip.IsFinished {
		return ErrCanNotUpdate
	}

	if newTrip.StartDate != nil {
		if !newTrip.StartDate.IsZero() {
			if oldTrip.SoldTickets != 0 {
				return ErrCanNotUpdate
			} else {
				updates["start_date"] = newTrip.StartDate
			}
		}
	}

	if newTrip.SoldTickets != 0 {
		updates["sold_tickets"] = newTrip.SoldTickets
	}

	err := s.repo.UpdateTrip(ctx, newTrip.Id, updates)
	if err != nil {
		log.Println("error on updating trip: ", err.Error())
		return err
	}

	return nil
}

func (s *service) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteTrip(ctx, id)
	if err != nil {
		log.Println("error on deleting trip: ", err.Error())
		return err
	}

	return nil
}

func (s *service) GetTrips(ctx context.Context, pageSize int, pageNumber int) ([]domain.Trip, error) {
	trips, err := s.repo.GetTrips(ctx, pageSize, pageNumber)
	if err != nil {
		log.Println("error on getting trips: ", err.Error())
		return nil, err
	}

	return trips, nil
}

func (s *service) ConfirmTrip(ctx context.Context, id uuid.UUID, userId uuid.UUID) error {
	trip, err := s.tripRepo.GetTripById(ctx, id)
	if err != nil {
		return fmt.Errorf("error on confirm trip: %s", err.Error())
	}
	if trip.TechnicalTeamID == nil {
		return fmt.Errorf("failed to confirm trip it dose not have technical team")
	}
	isTechnicalTeamMember, err := s.technicalTeamRepo.IsUserTechnicalTeamMember(ctx, *trip.TechnicalTeamID, userId)
	if err != nil {
		return fmt.Errorf("error on confirm trip: %s", err.Error())
	}
	if !isTechnicalTeamMember {
		return fmt.Errorf("failed to confirm trip you are not technical team member")
	}
	err = s.repo.ConfirmTrip(ctx, id)
	if err != nil {
		return fmt.Errorf("error on confirm trip: %s", err.Error())
	}

	return nil
}
