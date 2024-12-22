package domain

import (
	"time"

	"github.com/google/uuid"
)

// FactorUUID is a UUID used for external references.
type FactorUUID = uuid.UUID

// FactorStatus defines the status of a factor.
type FactorStatus uint8

const (
	FactorStatusUnknown FactorStatus = iota
	FactorStatusPending
	FactorStatusApproved
	FactorStatusRejected
)

// Factor represents an invoice or payment request in the system.
type Factor struct {
	ID            FactorUUID
	SourceService string       // The service generating the factor (e.g., HotelService, TransportService).
	ExternalID    string       // Unique identifier in the source service.
	BookingID     string       // ID linking the factor to a booking.
	Amount        uint         // The total amount for the factor.
	CustomerID    uint         // Reference to the customer associated with the factor.
	Status        FactorStatus // Current status of the factor.
	Details       string       // JSON or additional metadata related to the factor.
	CreatedAt     time.Time    // Timestamp when the factor was created.
	UpdatedAt     time.Time    // Timestamp when the factor was last updated.
}

// FactorFilters are used to filter factor queries.
type FactorFilters struct {
	SourceService string       // Filter by source service.
	BookingID     string       // Filter by booking ID.
	CustomerID    uuid.UUID    // Filter by customer ID.
	Status        FactorStatus // Filter by factor status.
}
