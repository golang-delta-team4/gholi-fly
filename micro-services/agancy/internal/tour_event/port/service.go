package port

import (
	"context"
	"gholi-fly-agancy/internal/tour_event/domain"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

// TourEventService defines the interface for TourEvent service layer.
type TourEventService interface {
	// Create a new TourEvent
	CreateEvent(ctx context.Context, events []domain.TourEvent) error

	// Update an existing TourEvent
	UpdateEvent(ctx context.Context, event domain.TourEvent) error

	// Get a single TourEvent by ID
	GetEventByID(ctx context.Context, id uuid.UUID) (*domain.TourEvent, error)

	// Search TourEvents based on filters, sort, and limit
	SearchEvents(ctx context.Context, search domain.TourEventSearch) ([]domain.TourEvent, error)

	// Get all TourEvents related to a specific ReservationID
	GetEventsByReservationID(ctx context.Context, reservationID uuid.UUID) ([]domain.TourEvent, error)

	// Delete a TourEvent by ID
	DeleteEvent(ctx context.Context, id uuid.UUID) error

	RegisterSagaRunner(scheduler gocron.Scheduler)
}
