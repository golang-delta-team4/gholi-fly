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
