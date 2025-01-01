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

	var technicalTeamId *uuid.UUID
	if req.TechnicalTeamId != "" {
		technicalTeamIdTemp, err := uuid.Parse(req.TechnicalTeamId)
		if err != nil {
			return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
		technicalTeamId = &technicalTeamIdTemp
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
		CompanyID:                companyId,
		TripType:                 req.TripType,
		UserReleaseDate:          userReleaseDate,
		TourReleaseDate:          tourReleaseDate,
		UserPrice:                float64(req.UserPrice),
		AgencyPrice:              float64(req.AgencyPrice),
		PathID:                   pathId,
		MinPassengers:            uint(req.MinPassengers),
		TechnicalTeamID:          technicalTeamId,
		VehicleYearOfManufacture: int(req.VehicleYearOfManufacture),
		SoldTickets:              uint(req.SoldTickets),
		MaxTickets:               uint(req.MaxTickets),
		StartDate:                &startDate,
		EndDate:                  &endDate,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateTripResponse{
		Id: tripId.String(),
	}, nil
}

func (s *TripService) GetTripById(ctx context.Context, tripId string) (*pb.GetTripResponse, error) {
	tripUId, err := uuid.Parse(tripId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	trip, err := s.svc.GetTripById(ctx, tripUId)

	if err != nil {
		return nil, err
	}

	var technicalTeamId string
	if trip.TechnicalTeamID != nil {
		technicalTeamId = trip.TechnicalTeamID.String()
	}

	return &pb.GetTripResponse{
		Id:               trip.Id.String(),
		CompanyId:        trip.CompanyID.String(),
		TripType:         trip.TripType,
		ReleaseDate:      trip.UserReleaseDate.String(),
		Price:            trip.UserPrice,
		PathId:           trip.PathID.String(),
		FromCountry:      trip.FromCountry,
		ToCountry:        trip.ToCountry,
		Origin:           trip.Origin,
		FromTerminalName: trip.FromTerminalName,
		ToTerminalName:   trip.ToTerminalName,
		Destination:      trip.Destination,
		PathName:         trip.PathName,
		PathDistanceKm:   trip.PathDistanceKM,
		Status:           trip.Status,
		MinPassengers:    uint32(trip.MinPassengers),
		TechnicalTeamId:  technicalTeamId,
		StartDate:        trip.StartDate.String(),
		EndDate:          trip.EndDate.String(),
	}, nil
}

func (s *TripService) GetAgencyTripById(ctx context.Context, tripId string) (*pb.GetTripResponse, error) {
	tripUId, err := uuid.Parse(tripId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	trip, err := s.svc.GetTripById(ctx, tripUId)

	if err != nil {
		return nil, err
	}
	var technicalTeamId string
	if trip.TechnicalTeamID != nil {
		technicalTeamId = trip.TechnicalTeamID.String()
	}
	return &pb.GetTripResponse{
		Id:               trip.Id.String(),
		CompanyId:        trip.CompanyID.String(),
		TripType:         trip.TripType,
		ReleaseDate:      trip.TourReleaseDate.String(),
		Price:            trip.AgencyPrice,
		PathId:           trip.PathID.String(),
		FromCountry:      trip.FromCountry,
		ToCountry:        trip.ToCountry,
		Origin:           trip.Origin,
		FromTerminalName: trip.FromTerminalName,
		ToTerminalName:   trip.ToTerminalName,
		Destination:      trip.Destination,
		PathName:         trip.PathName,
		PathDistanceKm:   trip.PathDistanceKM,
		Status:           trip.Status,
		MinPassengers:    uint32(trip.MinPassengers),
		TechnicalTeamId:  technicalTeamId,
		StartDate:        trip.StartDate.String(),
		EndDate:          trip.EndDate.String(),
	}, nil
}

func (s *TripService) GetTrips(ctx context.Context, pageSize int, pageNumber int) (*pb.GetTripsResponse, error) {
	trips, err := s.svc.GetTrips(ctx, pageSize, pageNumber)
	if err != nil {
		return nil, err
	}

	var response []*pb.GetTripResponse
	for _, trip := range trips {
		var technicalTeamId string
		if trip.TechnicalTeamID != nil {
			technicalTeamId = trip.TechnicalTeamID.String()
		}
		response = append(response, &pb.GetTripResponse{
			Id:               trip.Id.String(),
			CompanyId:        trip.CompanyID.String(),
			TripType:         trip.TripType,
			ReleaseDate:      trip.UserReleaseDate.String(),
			Price:            trip.UserPrice,
			PathId:           trip.PathID.String(),
			FromCountry:      trip.FromCountry,
			ToCountry:        trip.ToCountry,
			Origin:           trip.Origin,
			FromTerminalName: trip.FromTerminalName,
			ToTerminalName:   trip.ToTerminalName,
			Destination:      trip.Destination,
			PathName:         trip.PathName,
			PathDistanceKm:   trip.PathDistanceKM,
			Status:           trip.Status,
			MinPassengers:    uint32(trip.MinPassengers),
			TechnicalTeamId:  technicalTeamId,
			StartDate:        trip.StartDate.String(),
			EndDate:          trip.EndDate.String(),
		})
	}
	return &pb.GetTripsResponse{
		Trips: response,
	}, nil
}

func (s *TripService) GetAgencyTrips(ctx context.Context, pageSize int, pageNumber int) (*pb.GetTripsResponse, error) {
	trips, err := s.svc.GetTrips(ctx, pageSize, pageNumber)
	if err != nil {
		return nil, err
	}

	var response []*pb.GetTripResponse
	for _, trip := range trips {
		var technicalTeamId string
		if trip.TechnicalTeamID != nil {
			technicalTeamId = trip.TechnicalTeamID.String()
		}
		response = append(response, &pb.GetTripResponse{
			Id:               trip.Id.String(),
			CompanyId:        trip.CompanyID.String(),
			TripType:         trip.TripType,
			ReleaseDate:      trip.TourReleaseDate.String(),
			Price:            trip.AgencyPrice,
			PathId:           trip.PathID.String(),
			FromCountry:      trip.FromCountry,
			ToCountry:        trip.ToCountry,
			Origin:           trip.Origin,
			FromTerminalName: trip.FromTerminalName,
			ToTerminalName:   trip.ToTerminalName,
			Destination:      trip.Destination,
			PathName:         trip.PathName,
			PathDistanceKm:   trip.PathDistanceKM,
			Status:           trip.Status,
			MinPassengers:    uint32(trip.MinPassengers),
			TechnicalTeamId:  technicalTeamId,
			StartDate:        trip.StartDate.String(),
			EndDate:          trip.EndDate.String(),
		})
	}
	return &pb.GetTripsResponse{
		Trips: response,
	}, nil
}

func (s *TripService) UpdateTrip(ctx context.Context, tripId string, req *pb.UpdateTripRequest) error {

	tripUId, err := uuid.Parse(tripId)
	if err != nil {
		return fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}

	userReleaseDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.UserReleaseDate)
	if err != nil {
		return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	tourReleaseDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.TourReleaseDate)
	if err != nil {
		return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	var technicalTeamId *uuid.UUID
	if req.TechnicalTeamId != "" {
		technicalTeamIdTemp, err := uuid.Parse(req.TechnicalTeamId)
		if err != nil {
			return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
		technicalTeamId = &technicalTeamIdTemp
	}

	var vehicleRequestID *uuid.UUID
	if req.VehicleRequestId != "" {
		vehicleRequestIDTemp, err := uuid.Parse(req.VehicleRequestId)
		if err != nil {
			return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
		vehicleRequestID = &vehicleRequestIDTemp
	}

	var startDate *time.Time
	if req.StartDate != "" {
		tempStartDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.StartDate)
		if err != nil {
			return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
		startDate = &tempStartDate
	}

	var endDate *time.Time
	if req.EndDate != "" {
		tempEndDate, err := time.Parse("2006-01-02 15:04:05.999999-07:00", req.EndDate)
		if err != nil {
			return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
		}
		endDate = &tempEndDate
	}

	oldTrip, err := s.svc.GetTripById(ctx, tripUId)
	if err != nil {
		return err
	}

	err = s.svc.UpdateTrip(ctx, domain.Trip{
		Id:               tripUId,
		TripType:         req.TripType,
		UserReleaseDate:  userReleaseDate,
		TourReleaseDate:  tourReleaseDate,
		UserPrice:        float64(req.UserPrice),
		AgencyPrice:      float64(req.AgencyPrice),
		MinPassengers:    uint(req.MinPassengers),
		TechnicalTeamID:  technicalTeamId,
		VehicleRequestID: vehicleRequestID,
		SoldTickets:      uint(req.SoldTickets),
		MaxTickets:       uint(req.MaxTickets),
		StartDate:        startDate,
		EndDate:          endDate,
	}, *oldTrip)

	return err
}

func (s *TripService) DeleteTrip(ctx context.Context, tripId string) error {
	tripUId, err := uuid.Parse(tripId)
	if err != nil {
		return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	return s.svc.DeleteTrip(ctx, tripUId)
}

func (s *TripService) ConfirmTrip(ctx context.Context, tripId string, userId uuid.UUID) error {
	tripUId, err := uuid.Parse(tripId)
	if err != nil {
		return fmt.Errorf("%w %w", ErrTripCreationValidation, err)
	}

	return s.svc.ConfirmTrip(ctx, tripUId, userId)
}
