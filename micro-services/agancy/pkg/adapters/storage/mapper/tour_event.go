package mapper

import (
	"gholi-fly-agancy/internal/tour_event/domain"
	"gholi-fly-agancy/pkg/adapters/storage/types"
	"gholi-fly-agancy/pkg/fp"
)

// TourEventDomain2Storage converts a TourEvent from the domain layer to the storage layer.
func TourEventDomain2Storage(eventDomain domain.TourEvent) *types.TourEvent {
	return &types.TourEvent{
		ID:                  eventDomain.ID,
		ReservationID:       eventDomain.ReservationID,
		EventType:           eventDomain.EventType,
		Payload:             eventDomain.Payload,
		CompensationPayload: eventDomain.CompensationPayload,
		Status:              eventDomain.Status,
		RetryCount:          eventDomain.RetryCount,
		CreatedAt:           eventDomain.CreatedAt,
		UpdatedAt:           eventDomain.UpdatedAt,
	}
}

func tourEventDomain2Storage(eventDomain domain.TourEvent) types.TourEvent {
	return types.TourEvent{
		ID:                  eventDomain.ID,
		ReservationID:       eventDomain.ReservationID,
		EventType:           eventDomain.EventType,
		Payload:             eventDomain.Payload,
		CompensationPayload: eventDomain.CompensationPayload,
		Status:              eventDomain.Status,
		RetryCount:          eventDomain.RetryCount,
		CreatedAt:           eventDomain.CreatedAt,
		UpdatedAt:           eventDomain.UpdatedAt,
	}
}

func BatchTourEventDomain2Storage(domains []domain.TourEvent) []types.TourEvent {
	return fp.Map(domains, tourEventDomain2Storage)
}

// TourEventStorage2Domain converts a TourEvent from the storage layer to the domain layer.
func TourEventStorage2Domain(event types.TourEvent) *domain.TourEvent {
	return &domain.TourEvent{
		ID:                  event.ID,
		ReservationID:       event.ReservationID,
		EventType:           domain.EventType(event.EventType),
		Payload:             event.Payload,
		CompensationPayload: event.CompensationPayload,
		Status:              domain.EventStatus(event.Status),
		RetryCount:          event.RetryCount,
		CreatedAt:           event.CreatedAt,
		UpdatedAt:           event.UpdatedAt,
	}
}

func tourEventStorage2Domain(event types.TourEvent) domain.TourEvent {
	return domain.TourEvent{
		ID:                  event.ID,
		ReservationID:       event.ReservationID,
		EventType:           domain.EventType(event.EventType),
		Payload:             event.Payload,
		CompensationPayload: event.CompensationPayload,
		Status:              domain.EventStatus(event.Status),
		RetryCount:          event.RetryCount,
		CreatedAt:           event.CreatedAt,
		UpdatedAt:           event.UpdatedAt,
	}
}

func BatchTourEventStorage2Domain(events []types.TourEvent) []domain.TourEvent {
	return fp.Map(events, tourEventStorage2Domain)
}
