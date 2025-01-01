package port

import (
	"context"
	"gholi-fly-agancy/internal/tour_event/domain"

	"github.com/google/uuid"
)

// TourEventRepo defines the interface for interacting with TourEvent storage
type TourEventRepo interface {
	// Create a new event
	Create(ctx context.Context, events []domain.TourEvent) error

	// Update an existing event
	Update(ctx context.Context, event domain.TourEvent) error

	// Get a single event by ID
	GetByID(ctx context.Context, id uuid.UUID) (*domain.TourEvent, error)

	// Search events with filters, sort, and limit (for poller or saga workflows)
	Search(ctx context.Context, search domain.TourEventSearch) ([]domain.TourEvent, error)

	// Get all events related to a specific ReservationID
	GetByReservationID(ctx context.Context, reservationID uuid.UUID) ([]domain.TourEvent, error)

	// Delete an event by ID
	Delete(ctx context.Context, id uuid.UUID) error
}
