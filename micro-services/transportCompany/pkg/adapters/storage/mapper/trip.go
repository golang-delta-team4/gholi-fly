package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
)

func TripDomain2Storage(tripDomain domain.Trip) *types.Trip {
	return &types.Trip{
		Id:               tripDomain.Id,
		CompanyID:        tripDomain.CompanyID,
		UserReleaseDate:  tripDomain.UserReleaseDate,
		TourReleaseDate:  tripDomain.TourReleaseDate,
		TripType:         tripDomain.TripType,
		UserPrice:        tripDomain.UserPrice,
		AgencyPrice:      tripDomain.AgencyPrice,
		PathID:           tripDomain.PathID,
		MinPassengers:    tripDomain.MinPassengers,
		FromCountry:      tripDomain.FromCountry,
		ToCountry:        tripDomain.ToCountry,
		FromTerminalName: tripDomain.FromTerminalName,
		Origin:           tripDomain.Origin,
		SoldTickets:      tripDomain.SoldTickets,
		ToTerminalName:   tripDomain.ToTerminalName,
		Destination:      tripDomain.Destination,
		PathName:         tripDomain.PathName,
		PathDistanceKM:   tripDomain.PathDistanceKM,
		IsCanceled:       tripDomain.IsCanceled,
		IsConfirmed:      tripDomain.IsConfirmed,
		Status:           tripDomain.Status,
		StartDate:        tripDomain.StartDate,
		EndDate:          tripDomain.EndDate,
		Profit:           tripDomain.Profit,
		MaxTickets:       tripDomain.MaxTickets,
		TechnicalTeamID:  tripDomain.TechnicalTeamID,
		VehicleRequestID: tripDomain.VehicleRequestID,
		VehicleID:        tripDomain.VehicleID,
		IsFinished:       tripDomain.IsFinished,
		VehicleName:      tripDomain.VehicleName,
	}
}

func TripStorage2Domain(tripStorage types.Trip) *domain.Trip {
	return &domain.Trip{
		Id:               tripStorage.Id,
		CompanyID:        tripStorage.CompanyID,
		UserReleaseDate:  tripStorage.UserReleaseDate,
		TourReleaseDate:  tripStorage.TourReleaseDate,
		TripType:         tripStorage.TripType,
		UserPrice:        tripStorage.UserPrice,
		AgencyPrice:      tripStorage.AgencyPrice,
		PathID:           tripStorage.PathID,
		MinPassengers:    tripStorage.MinPassengers,
		FromCountry:      tripStorage.FromCountry,
		ToCountry:        tripStorage.ToCountry,
		FromTerminalName: tripStorage.FromTerminalName,
		Origin:           tripStorage.Origin,
		SoldTickets:      tripStorage.SoldTickets,
		ToTerminalName:   tripStorage.ToTerminalName,
		Destination:      tripStorage.Destination,
		PathName:         tripStorage.PathName,
		PathDistanceKM:   tripStorage.PathDistanceKM,
		IsCanceled:       tripStorage.IsCanceled,
		IsConfirmed:      tripStorage.IsConfirmed,
		Status:           tripStorage.Status,
		StartDate:        tripStorage.StartDate,
		EndDate:          tripStorage.EndDate,
		Profit:           tripStorage.Profit,
		MaxTickets:       tripStorage.MaxTickets,
		TechnicalTeamID:  tripStorage.TechnicalTeamID,
		VehicleRequestID: tripStorage.VehicleRequestID,
		VehicleID:        tripStorage.VehicleID,
		IsFinished:       tripStorage.IsFinished,
		VehicleName:      tripStorage.VehicleName,
	}
}
