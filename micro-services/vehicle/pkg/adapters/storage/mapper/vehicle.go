package mapper

import (
	"vehicle/internal/vehicle/domain"
	"vehicle/pkg/adapters/storage/types"
)

func VehicleToDomain(v *types.Vehicle) *domain.Vehicle {
	return &domain.Vehicle{
		ID:         v.ID,
		OwnerID:    v.OwnerID,
		Type:       v.Type,
		Capacity:   v.Capacity,
		Speed:      v.Speed,
		UniqueCode: v.UniqueCode,
		Status:     v.Status,
		CreatedAt:  v.CreatedAt,
		UpdatedAt:  v.UpdatedAt,
	}
}

func DomainToVehicle(d *domain.Vehicle) *types.Vehicle {
	return &types.Vehicle{
		ID:         d.ID,
		OwnerID:    d.OwnerID,
		Type:       d.Type,
		Capacity:   d.Capacity,
		Speed:      d.Speed,
		UniqueCode: d.UniqueCode,
		Status:     d.Status,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}
