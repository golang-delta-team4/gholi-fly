package vehicle

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
	"vehicle/internal/vehicle/domain"
	"vehicle/internal/vehicle/port"
	"vehicle/pkg/adapters/transport"

	"github.com/google/uuid"
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
	return s.repo.GetByID(ctx, id)
}

func (s *service) GetAllVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) UpdateVehicle(ctx context.Context, vehicle *domain.Vehicle) error {
	return s.repo.Update(ctx, vehicle)
}

func (s *service) DeleteVehicle(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) MatchVehicle(ctx context.Context, tripRequest *domain.TripRequest) (*domain.Vehicle, error) {
	vehicles, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	filteredVehicles := []domain.Vehicle{}
	for _, vehicle := range vehicles {
		if vehicle.Type == tripRequest.TripType && vehicle.Capacity >= tripRequest.MinPassengers {
			filteredVehicles = append(filteredVehicles, vehicle)
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, fmt.Errorf("no matching vehicles found")
	}

	sort.Slice(filteredVehicles, func(i, j int) bool {
		if filteredVehicles[i].Capacity != filteredVehicles[j].Capacity {
			return filteredVehicles[i].Capacity > filteredVehicles[j].Capacity
		}
		if filteredVehicles[i].YearOfManufacture != filteredVehicles[j].YearOfManufacture {
			return filteredVehicles[i].YearOfManufacture > filteredVehicles[j].YearOfManufacture
		}
		return filteredVehicles[i].CreatedAt.Before(filteredVehicles[j].CreatedAt)
	})

	return &filteredVehicles[0], nil
}
