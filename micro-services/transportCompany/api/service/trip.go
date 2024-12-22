package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	tripPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/google/uuid"
)

type TripService struct {
	svc tripPort.Service
}

func NewTripService(svc tripPort.Service) *TripService {
	return &TripService{
		svc: svc,
	}
}

var (
	ErrTripCreationValidation = trip.ErrTripCreationValidation
)

func (s *TripService) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {

	companyId, err := uuid.Parse(req.CompanyId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	pathId, err := uuid.Parse(req.PathId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	userReleaseDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.UserReleaseDate)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	tourReleaseDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.TourReleaseDate)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	var technicalTeamId uuid.UUID
	if req.TechnicalTeamId != "" {
		technicalTeamId, err = uuid.Parse(req.TechnicalTeamId)
		if err != nil {
			return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
	}

	var vehicleRequestID uuid.UUID
	if req.VehicleRequestId != "" {
		vehicleRequestID, err = uuid.Parse(req.VehicleRequestId)
		if err != nil {
			return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
	}

	startDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	endDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	tripId, err := s.svc.CreateTrip(ctx, domain.Trip{
		CompanyID:        companyId,
		TripType:         req.TripType,
		UserReleaseDate:  userReleaseDate,
		TourReleaseDate:  tourReleaseDate,
		UserPrice:        float64(req.UserPrice),
		AgencyPrice:      float64(req.AgencyPrice),
		PathID:           pathId,
		MinPassengers:    uint(req.MinPassengers),
		TechnicalTeamID:  &technicalTeamId,
		VehicleRequestID: &vehicleRequestID,
		SoldTickets:      uint(req.SoldTickets),
		MaxTickets:       uint(req.MaxTickets),
		StartDate:        &startDate,
		EndDate:          &endDate,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateTripResponse{
		Id: tripId.String(),
	}, nil
}
