package presenter

import "github.com/google/uuid"

type GetPathByIDResponse struct {
	SourceTerminalID      uuid.UUID `json:"source_terminal_id"`
	DestinationTerminalID uuid.UUID `json:"destination_terminal_id"`
	DistanceKM            float64   `json:"distance_km"`
	RouteCode             string    `json:"route_code"`
	VehicleType           string    `json:"vehicle_type"`
	SourceTerminal        *Terminal `json:"source_terminal"`
	DestinationTerminal   *Terminal `json:"destination_terminal"`
}
