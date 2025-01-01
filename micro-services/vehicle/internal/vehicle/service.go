package vehicle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"vehicle/internal/vehicle/domain"
	"vehicle/internal/vehicle/port"
	"vehicle/pkg/adapters/storage/mapper"
	"vehicle/pkg/adapters/storage/types"
	"vehicle/pkg/adapters/transport"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type service struct {
	repo           port.VehicleRepository
	tripServiceURL string
}

func NewVehicleService(repo port.VehicleRepository, tripServiceURL string) port.VehicleService {
	return &service{
		repo:           repo,
		tripServiceURL: tripServiceURL,
	}
}

var (
	ErrVehicleNotFound = errors.New("vehicle not found")
)

// Example of using the tripServiceURL in one of the methods
func (s *service) FetchTripRequest(ctx context.Context) (*domain.TripRequest, error) {
	return transport.GetTripRequest(s.tripServiceURL)
}

func (s *service) ProcessTripRequest(ctx context.Context) (*domain.TripRequest, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.tripServiceURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch trip request: %s", resp.Status)
	}

	var tripRequest domain.TripRequest
	if err := json.NewDecoder(resp.Body).Decode(&tripRequest); err != nil {
		return nil, err
	}

	return &tripRequest, nil
}

// Implement the methods defined in the VehicleService interface.
func (s *service) CreateVehicle(ctx context.Context, vehicle *domain.Vehicle) error {
	// Validate UniqueCode
	if vehicle.UniqueCode == "" {
		return fmt.Errorf("unique code is required")
	}

	// Validate YearOfManufacture
	if vehicle.YearOfManufacture < 1900 || vehicle.YearOfManufacture > time.Now().Year() {
		return fmt.Errorf("year of manufacture must be between 1900 and the current year")
	}

	return s.repo.Create(ctx, vehicle)
}

func (s *service) GetVehicleByID(ctx context.Context, id uuid.UUID) (*domain.Vehicle, error) {
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.VehicleToDomain(vehicle), nil
}

func (s *service) GetAllVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) UpdateVehicle(ctx context.Context, vehicle *domain.Vehicle) error {
	currentVehicle, err := s.repo.GetByID(ctx, vehicle.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrVehicleNotFound
		}
		return err
	}
	if vehicle.Capacity != 0 {
		currentVehicle.Capacity = vehicle.Capacity
	}
	if vehicle.PricePerKilometer != 0 {
		currentVehicle.PricePerKilometer = vehicle.PricePerKilometer
	}
	if vehicle.Speed != 0 {
		currentVehicle.Speed = vehicle.Speed
	}
	if vehicle.Status != "" {
		currentVehicle.Status = vehicle.Status
	}
	return s.repo.Update(ctx, currentVehicle)
}

func (s *service) DeleteVehicle(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) MatchVehicle(ctx context.Context, vehicleMatchRequest *domain.MatchMakerRequest) (uuid.UUID, *domain.Vehicle, error) {
	vehicle, err := s.repo.GetMatchedVehicle(ctx, vehicleMatchRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uuid.Nil, nil, ErrVehicleNotFound
		}
		return uuid.Nil, nil, err
	}
	vehicleDomain := mapper.VehicleToDomain(&vehicle)
	reservationID, err := s.repo.CreateReservation(ctx, types.VehicleReserve{
		TripID:    vehicleMatchRequest.TripID,
		StartDate: vehicleMatchRequest.ReserveStartDate,
		EndDate:   vehicleMatchRequest.ReserveEndDate,
		VehicleID: vehicle.ID})
	if err != nil {
		return uuid.Nil, nil, err
	}
	return reservationID, vehicleDomain, nil
}
