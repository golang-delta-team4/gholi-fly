package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vehicle/internal/vehicle/domain"
)

func GetTripRequest(url string) (*domain.TripRequest, error) {
	// Create an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch trip request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the JSON response
	var tripRequest domain.TripRequest
	if err := json.NewDecoder(resp.Body).Decode(&tripRequest); err != nil {
		return nil, fmt.Errorf("failed to decode trip request: %w", err)
	}

	return &tripRequest, nil
}
