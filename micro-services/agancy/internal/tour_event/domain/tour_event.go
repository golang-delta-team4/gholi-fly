package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EventType string
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte but got %T", value)
	}
	if err := json.Unmarshal(bytes, j); err != nil {
		return fmt.Errorf("failed to unmarshal JSONB: %w", err)
	}
	return nil
}

const (
	EventTypeHotelReservation EventType = "hotel_reservation"
	EventTypeTripReservation  EventType = "trip_reservation"
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type EventStatus string

const (
	StatusPending             EventStatus = "pending"
	StatusSuccess             EventStatus = "success"
	StatusFailed              EventStatus = "failed"
	StatusCompensationPending EventStatus = "compensation_pending"
	StatusCompensated         EventStatus = "compensated"
)

type TourEvent struct {
	ID                  uuid.UUID   `json:"id"`
	ReservationID       uuid.UUID   `json:"reservation_id"`
	EventType           EventType   `json:"event_type"`
	Payload             string      `json:"payload"`
	CompensationPayload string      `json:"compensation_payload"`
	Status              EventStatus `json:"status"`
	RetryCount          int         `json:"retry_count"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
}

// TourEventSearch defines the query parameters for searching TourEvent records
type TourEventSearch struct {
	EventType   *EventType   `json:"event_type,omitempty"`     // Filter by EventType
	Status      *EventStatus `json:"status,omitempty"`         // Filter by Status
	Reservation *uuid.UUID   `json:"reservation_id,omitempty"` // Filter by ReservationID
	SortBy      string       `json:"sort_by,omitempty"`        // Sort field (e.g., "created_at")
	SortOrder   SortOrder    `json:"sort_order,omitempty"`     // Sort type (asc or desc)
	LimitCount  uint         `json:"limit_count,omitempty"`    // Maximum number of records to return
}

func (te *TourEvent) Validate() error {
	if te.ReservationID == uuid.Nil {
		return errors.New("invalid reservation ID")
	}
	if te.EventType == "" {
		return errors.New("event type is required")
	}
	if te.Status == "" {
		return errors.New("status is required")
	}
	return nil
}
