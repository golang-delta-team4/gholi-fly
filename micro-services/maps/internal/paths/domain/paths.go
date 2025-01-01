package domain

import (
	"time"

	"github.com/google/uuid"
)

type Path struct {
	ID                    uuid.UUID `json:"id"`
	SourceTerminalID      uuid.UUID `json:"source_terminal_id"`
	DestinationTerminalID uuid.UUID `json:"destination_terminal_id"`
	DistanceKM            float64   `json:"distance_km"`
	RouteCode             string    `json:"route_code"`
	VehicleType           string    `json:"vehicle_type"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
