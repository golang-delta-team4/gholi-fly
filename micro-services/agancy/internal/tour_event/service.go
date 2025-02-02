package tour_event

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"gholi-fly-agancy/internal/tour_event/domain"
	"gholi-fly-agancy/internal/tour_event/port"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

var (
	ErrEventCreateFailed   = errors.New("failed to create tour event")
	ErrEventNotFound       = errors.New("tour event not found")
	ErrEventValidation     = errors.New("invalid tour event details")
	MaxCompensationRetries = 5
)

type service struct {
	repo port.TourEventRepo
}

func NewService(repo port.TourEventRepo) port.TourEventService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateEvent(ctx context.Context, events []domain.TourEvent) error {
	for _, event := range events {
		if err := event.Validate(); err != nil {
			return fmt.Errorf("%w: %v", ErrEventValidation, err)
		}
	}
	err := s.repo.Create(ctx, events)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrEventCreateFailed, err)
	}

	return nil
}

func (s *service) UpdateEvent(ctx context.Context, event domain.TourEvent) error {
	if err := event.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrEventValidation, err)
	}

	if err := s.repo.Update(ctx, event); err != nil {
		return err
	}

	return nil
}

func (s *service) GetEventByID(ctx context.Context, id uuid.UUID) (*domain.TourEvent, error) {
	event, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, ErrEventNotFound
	}

	return event, nil
}

func (s *service) SearchEvents(ctx context.Context, search domain.TourEventSearch) ([]domain.TourEvent, error) {
	events, err := s.repo.Search(ctx, search)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) GetEventsByReservationID(ctx context.Context, reservationID uuid.UUID) ([]domain.TourEvent, error) {
	events, err := s.repo.GetByReservationID(ctx, reservationID)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

// RegisterSagaRunner initializes and registers the saga runner.
func (s *service) RegisterSagaRunner(scheduler gocron.Scheduler) {
	// Schedule the compensation logic to run every 10 seconds.
	_, err := scheduler.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(func() {
			ctx := context.Background()

			// Query failed events
			failedStatus := domain.StatusFailed
			failedEvents, err := s.repo.Search(ctx, domain.TourEventSearch{
				Status:     &failedStatus,
				SortBy:     "created_at",
				SortOrder:  domain.SortOrderAsc,
				LimitCount: 10,
			})
			if err != nil {
				log.Printf("Failed to fetch failed events: %v\n", err)
				return
			}

			for _, event := range failedEvents {
				relatedEvents, err := s.repo.GetByReservationID(ctx, event.ReservationID)
				if err != nil {
					log.Printf("Failed to fetch related events for reservation %s: %v\n", event.ReservationID, err)
					continue
				}

				successfulRelatedEvents := filterSuccessfulEvents(relatedEvents)

				for _, relatedEvent := range successfulRelatedEvents {
					err := s.compensate(relatedEvent)
					if err != nil {
						log.Printf("Failed to compensate event %s: %v\n", relatedEvent.ID, err)

						if relatedEvent.RetryCount < MaxCompensationRetries {
							relatedEvent.RetryCount++
							err = s.repo.Update(ctx, relatedEvent)
							if err != nil {
								log.Printf("Failed to update retry count for event %s: %v\n", relatedEvent.ID, err)
							}
						} else {
							log.Printf("Maximum retry reached for event %s; skipping compensation.\n", relatedEvent.ID)
						}
						continue
					}

					relatedEvent.Status = domain.StatusCompensated
					err = s.repo.Update(ctx, relatedEvent)
					if err != nil {
						log.Printf("Failed to update status for compensated event %s: %v\n", relatedEvent.ID, err)
					}
				}
			}
		}),
	)
	if err != nil {
		log.Println(err.Error())
	}

}

// compensate implements the compensation logic for a tour event.
func (s *service) compensate(event domain.TourEvent) error {
	switch event.EventType {
	case domain.EventTypeHotelReservation:
		fmt.Println(event.ReservationID)

		hotelURL := fmt.Sprintf("http://%s:%d/api/v1/hotel/booking/cancel/%s", "localhost", 8081, event.CompensationPayload)
		err := makePostRequest(context.Background(), hotelURL, nil, nil, "PATCH")
		if err != nil {
			return fmt.Errorf("failed to buy transport ticket: %w", err)
		}
		event.Status = domain.StatusCompensated
		s.repo.Update(context.Background(), event)
	case domain.EventTypeTripReservation:
		// Logic to compensate trip reservation
		transportURL := fmt.Sprintf("http://%s:%d/api/v1/transport-company/agency-ticket/cancel/%s", "localhost", 8085, event.CompensationPayload)
		err := makePostRequest(context.Background(), transportURL, nil, nil, "POST")
		if err != nil {
			return fmt.Errorf("failed to buy transport ticket: %w", err)
		}
		event.Status = domain.StatusCompensated
		s.repo.Update(context.Background(), event)
		log.Printf("Compensating trip reservation for event %s\n", event.ID)
	default:
		return fmt.Errorf("unknown event type: %s", event.EventType)
	}

	return nil
}

// filterSuccessfulEvents filters events with a "success" status.
func filterSuccessfulEvents(events []domain.TourEvent) []domain.TourEvent {
	successfulEvents := []domain.TourEvent{}
	for _, event := range events {
		if event.Status == domain.StatusSuccess {
			successfulEvents = append(successfulEvents, event)
		}
	}
	return successfulEvents
}
func makePostRequest(ctx context.Context, url string, requestBody interface{}, responseBody interface{}, method string) error {
	reqBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP request failed with status: %d", resp.StatusCode)
	}

	return nil
}
