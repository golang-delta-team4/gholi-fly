package mapper

import (
	"gholi-fly-agancy/internal/tour/domain"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"gholi-fly-agancy/pkg/fp"

	"github.com/google/uuid"
)

// tourDomain2Storage converts a tour from the domain layer to the storage layer.
func TourDomain2Storage(tourDomain domain.Tour) *types.Tour {
	return &types.Tour{
		ID:                  uuid.UUID(tourDomain.ID),
		Name:                tourDomain.Name,
		Description:         tourDomain.Description,
		StartDate:           tourDomain.StartDate,
		EndDate:             tourDomain.EndDate,
		SourceLocation:      tourDomain.SourceLocation,
		DestinationLocation: tourDomain.DestinationLocation,
		TripID:              tourDomain.TripID,
		TripAgencyPrice:     tourDomain.TripAgencyPrice,
		HotelID:             tourDomain.HotelID,
		HotelAgencyPrice:    tourDomain.HotelAgencyPrice,
		Capacity:            tourDomain.Capacity,
		IsPublished:         tourDomain.IsPublished,
		CreatedAt:           tourDomain.CreatedAt,
		UpdatedAt:           tourDomain.UpdatedAt,
	}
}

func tourDomain2Storage(tourDomain domain.Tour) types.Tour {
	return types.Tour{
		ID:                  uuid.UUID(tourDomain.ID),
		Name:                tourDomain.Name,
		Description:         tourDomain.Description,
		StartDate:           tourDomain.StartDate,
		EndDate:             tourDomain.EndDate,
		SourceLocation:      tourDomain.SourceLocation,
		DestinationLocation: tourDomain.DestinationLocation,
		TripID:              tourDomain.TripID,
		TripAgencyPrice:     tourDomain.TripAgencyPrice,
		HotelID:             tourDomain.HotelID,
		HotelAgencyPrice:    tourDomain.HotelAgencyPrice,
		Capacity:            tourDomain.Capacity,
		IsPublished:         tourDomain.IsPublished,
		CreatedAt:           tourDomain.CreatedAt,
		UpdatedAt:           tourDomain.UpdatedAt,
	}
}

func BatchTourDomain2Storage(domains []domain.Tour) []types.Tour {
	return fp.Map(domains, tourDomain2Storage)
}

// tourStorage2Domain converts a tour from the storage layer to the domain layer.
func TourStorage2Domain(tour types.Tour) *domain.Tour {
	return &domain.Tour{
		ID:                  domain.TourID(tour.ID),
		Name:                tour.Name,
		Description:         tour.Description,
		StartDate:           tour.StartDate,
		EndDate:             tour.EndDate,
		SourceLocation:      tour.SourceLocation,
		DestinationLocation: tour.DestinationLocation,
		TripID:              tour.TripID,
		TripAgencyPrice:     tour.TripAgencyPrice,
		HotelID:             tour.HotelID,
		HotelAgencyPrice:    tour.HotelAgencyPrice,
		Capacity:            tour.Capacity,
		IsPublished:         tour.IsPublished,
		CreatedAt:           tour.CreatedAt,
		UpdatedAt:           tour.UpdatedAt,
	}
}

func tourStorage2Domain(tour types.Tour) domain.Tour {
	return domain.Tour{
		ID:                  domain.TourID(tour.ID),
		Name:                tour.Name,
		Description:         tour.Description,
		StartDate:           tour.StartDate,
		EndDate:             tour.EndDate,
		SourceLocation:      tour.SourceLocation,
		DestinationLocation: tour.DestinationLocation,
		TripID:              tour.TripID,
		TripAgencyPrice:     tour.TripAgencyPrice,
		HotelID:             tour.HotelID,
		HotelAgencyPrice:    tour.HotelAgencyPrice,
		Capacity:            tour.Capacity,
		IsPublished:         tour.IsPublished,
		CreatedAt:           tour.CreatedAt,
		UpdatedAt:           tour.UpdatedAt,
	}
}

func BatchTourStorage2Domain(tours []types.Tour) []domain.Tour {
	return fp.Map(tours, tourStorage2Domain)
}
