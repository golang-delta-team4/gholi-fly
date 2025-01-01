package mapper

import (
	"gholi-fly-maps/internal/paths/domain"
	"gholi-fly-maps/pkg/adapters/storage/types"
)

// PathToDomain converts a database model to a domain model.
func PathToDomain(p *types.Path) *domain.Path {
	pathDomain := &domain.Path{
		ID:                  p.ID,
		DistanceKM:          p.DistanceKM,
		RouteCode:           p.RouteCode,
		VehicleType:         p.VehicleType,
		CreatedAt:           p.CreatedAt,
		UpdatedAt:           p.UpdatedAt,
	}
	if p.SourceTerminal != nil {
		pathDomain.SourceTerminal = TerminalToDomain(p.SourceTerminal)
	}
	if p.DestinationTerminal != nil {
		pathDomain.DestinationTerminal = TerminalToDomain(p.DestinationTerminal)
	}
	return pathDomain
}

// DomainToPath converts a domain model to a database model.
func DomainToPath(d *domain.Path) *types.Path {
	return &types.Path{
		ID:                    d.ID,
		SourceTerminalID:      d.SourceTerminalID,
		DestinationTerminalID: d.DestinationTerminalID,
		DistanceKM:            d.DistanceKM,
		RouteCode:             d.RouteCode,
		VehicleType:           d.VehicleType,
		CreatedAt:             d.CreatedAt,
		UpdatedAt:             d.UpdatedAt,
	}
}
