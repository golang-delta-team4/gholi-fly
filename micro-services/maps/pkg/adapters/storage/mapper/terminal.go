package mapper

import (
	"gholi-fly-maps/internal/terminals/domain"
	"gholi-fly-maps/pkg/adapters/storage/types"
)

// TerminalToDomain converts a database model to a domain model.
func TerminalToDomain(t *types.Terminal) *domain.Terminal {
	return &domain.Terminal{
		ID:        t.ID,
		Name:      t.Name,
		Location:  t.Location,
		Type:      string(t.Type),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// DomainToTerminal converts a domain model to a database model.
func DomainToTerminal(d *domain.Terminal) *types.Terminal {
	return &types.Terminal{
		ID:        d.ID,
		Name:      d.Name,
		Location:  d.Location,
		Type:      types.TerminalType(d.Type),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// batch
