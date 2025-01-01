package storage

import (
	"context"
	"errors"
	"gholi-fly-agancy/internal/tour_event/domain"
	"gholi-fly-agancy/internal/tour_event/port"
	"gholi-fly-agancy/pkg/adapters/storage/mapper"
	"gholi-fly-agancy/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tourEventRepo struct {
	db *gorm.DB
}

func NewTourEventRepo(db *gorm.DB) port.TourEventRepo {
	return &tourEventRepo{db}
}

func (r *tourEventRepo) Create(ctx context.Context, eventDomains []domain.TourEvent) error {
	event := mapper.BatchTourEventDomain2Storage(eventDomains)
	err := r.db.Table("tour_events").WithContext(ctx).Create(event).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *tourEventRepo) Update(ctx context.Context, eventDomain domain.TourEvent) error {
	event := mapper.TourEventDomain2Storage(eventDomain)
	return r.db.Table("tour_events").WithContext(ctx).Save(event).Error
}

func (r *tourEventRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.TourEvent, error) {
	var event types.TourEvent
	err := r.db.Table("tour_events").WithContext(ctx).First(&event, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.TourEventStorage2Domain(event), nil
}

func (r *tourEventRepo) Search(ctx context.Context, search domain.TourEventSearch) ([]domain.TourEvent, error) {
	var events []types.TourEvent
	query := r.db.Table("tour_events").WithContext(ctx)

	// Apply filters
	if search.EventType != nil {
		query = query.Where("event_type = ?", *search.EventType)
	}
	if search.Status != nil {
		query = query.Where("status = ?", *search.Status)
	}
	if search.Reservation != nil {
		query = query.Where("reservation_id = ?", *search.Reservation)
	}

	// Apply sorting
	if search.SortBy != "" {
		sortOrder := "ASC"
		if search.SortOrder == domain.SortOrderDesc {
			sortOrder = "DESC"
		}
		query = query.Order(search.SortBy + " " + sortOrder)
	}

	// Apply limit
	if search.LimitCount > 0 {
		query = query.Limit(int(search.LimitCount))
	}

	err := query.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchTourEventStorage2Domain(events), nil
}

func (r *tourEventRepo) GetByReservationID(ctx context.Context, reservationID uuid.UUID) ([]domain.TourEvent, error) {
	var events []types.TourEvent
	err := r.db.Table("tour_events").WithContext(ctx).Where("reservation_id = ?", reservationID).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchTourEventStorage2Domain(events), nil
}

func (r *tourEventRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Table("tour_events").WithContext(ctx).Delete(&types.TourEvent{}, "id = ?", id).Error
}
