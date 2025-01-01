package mapper

import (
	"gholi-fly-maps/internal/paths/domain"
	"gholi-fly-maps/pkg/adapters/storage/types"
)

// PathToDomain converts a database model to a domain model.
func PathToDomain(p *types.Path) *domain.Path {
	return &domain.Path{
		ID:                    p.ID,
		SourceTerminalID:      p.SourceTerminalID,
		DestinationTerminalID: p.DestinationTerminalID,
		DistanceKM:            p.DistanceKM,
		RouteCode:             p.RouteCode,
		VehicleType:           p.VehicleType,
		CreatedAt:             p.CreatedAt,
		UpdatedAt:             p.UpdatedAt,
	}
}

// DomainToPath converts a domain model to a database model.
func DomainToPath(d *domain.Path) *types.Path {
	return &types.Path{
		// ID:                    d.ID,
		SourceTerminalID:      d.SourceTerminalID,
		DestinationTerminalID: d.DestinationTerminalID,
		DistanceKM:            d.DistanceKM,
		RouteCode:             d.RouteCode,
		VehicleType:           d.VehicleType,
		CreatedAt:             d.CreatedAt,
		UpdatedAt:             d.UpdatedAt,
	}
}
