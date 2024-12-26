package mapper

import (
	"gholi-fly-agancy/internal/reservation/domain"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"gholi-fly-agancy/pkg/fp"

	"github.com/google/uuid"
)

// ReservationDomain2Storage converts a Reservation from the domain layer to the storage layer.
func ReservationDomain2Storage(reservationDomain domain.Reservation) *types.Reservation {
	return &types.Reservation{
		ID:         uuid.UUID(reservationDomain.ID),
		CustomerID: reservationDomain.CustomerID,
		FactorID:   reservationDomain.FactorID,
		TourID:     reservationDomain.TourID,
		Status:     reservationDomain.Status,
		CreatedAt:  reservationDomain.CreatedAt,
		UpdatedAt:  reservationDomain.UpdatedAt,
	}
}

func reservationDomain2Storage(reservationDomain domain.Reservation) types.Reservation {
	return types.Reservation{
		ID:         uuid.UUID(reservationDomain.ID),
		CustomerID: reservationDomain.CustomerID,
		FactorID:   reservationDomain.FactorID,
		TourID:     reservationDomain.TourID,
		Status:     reservationDomain.Status,
		CreatedAt:  reservationDomain.CreatedAt,
		UpdatedAt:  reservationDomain.UpdatedAt,
	}
}

func BatchReservationDomain2Storage(domains []domain.Reservation) []types.Reservation {
	return fp.Map(domains, reservationDomain2Storage)
}

// ReservationStorage2Domain converts a Reservation from the storage layer to the domain layer.
func ReservationStorage2Domain(reservation types.Reservation) *domain.Reservation {
	return &domain.Reservation{
		ID:         domain.ReservationID(reservation.ID),
		CustomerID: reservation.CustomerID,
		FactorID:   reservation.FactorID,
		TourID:     reservation.TourID,
		Status:     reservation.Status,
		CreatedAt:  reservation.CreatedAt,
		UpdatedAt:  reservation.UpdatedAt,
	}
}

func reservationStorage2Domain(reservation types.Reservation) domain.Reservation {
	return domain.Reservation{
		ID:         domain.ReservationID(reservation.ID),
		CustomerID: reservation.CustomerID,
		FactorID:   reservation.FactorID,
		TourID:     reservation.TourID,
		Status:     reservation.Status,
		CreatedAt:  reservation.CreatedAt,
		UpdatedAt:  reservation.UpdatedAt,
	}
}

func BatchReservationStorage2Domain(reservations []types.Reservation) []domain.Reservation {
	return fp.Map(reservations, reservationStorage2Domain)
}
