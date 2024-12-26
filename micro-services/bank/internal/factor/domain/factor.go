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
// Factor represents an invoice or payment request in the system.
type Factor struct {
	ID             FactorUUID   // Unique identifier for the factor.
	SourceService  string       // The service generating the factor (e.g., HotelService, TransportService).
	ExternalID     uuid.UUID    // Unique identifier in the source service.
	BookingID      uuid.UUID    // ID linking the factor to a booking.
	Amount         uint         // The total amount for the factor.
	CustomerID     uuid.UUID    // Reference to the customer associated with the factor.
	Status         FactorStatus // Current status of the factor.
	Details        string       // JSON or additional metadata related to the factor.
	InstantPayment bool         // Whether the payment should be made instantly.
	IsPaid         bool         // Whether the factor has been paid.
	CreatedAt      time.Time    // Timestamp when the factor was created.
	UpdatedAt      time.Time    // Timestamp when the factor was last updated.
}

// FactorFilters are used to filter factor queries.
type FactorFilters struct {
	FactorID      uuid.UUID    // Filter by specific factor ID (optional).
	SourceService string       // Filter by source service (optional).
	BookingID     string       // Filter by booking ID (optional).
	CustomerID    uuid.UUID    // Filter by customer ID (optional).
	IsPaid        *bool        // Filter by payment status (optional).
	Status        FactorStatus // Filter by factor status (optional).
	Page          int          // Pagination: Page number.
	PageSize      int          // Pagination: Number of results per page.
}
