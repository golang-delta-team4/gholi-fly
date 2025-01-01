package storage

import (
	"context"
	"errors"

	"gholi-fly-agancy/internal/reservation/domain"
	"gholi-fly-agancy/internal/reservation/port"
	"gholi-fly-agancy/pkg/adapters/storage/mapper"
	"gholi-fly-agancy/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reservationRepo struct {
	db *gorm.DB
}

func NewReservationRepo(db *gorm.DB) port.ReservationRepo {
	return &reservationRepo{db}
}

func (r *reservationRepo) Create(ctx context.Context, reservationDomain domain.Reservation) (domain.ReservationID, error) {
	reservation := mapper.ReservationDomain2Storage(reservationDomain)
	return domain.ReservationID(reservation.ID), r.db.Table("reservations").WithContext(ctx).Create(reservation).Error
}

func (r *reservationRepo) GetByID(ctx context.Context, id domain.ReservationID) (*domain.Reservation, error) {
	var reservation types.Reservation
	err := r.db.Table("reservations").WithContext(ctx).First(&reservation, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ReservationStorage2Domain(reservation), nil
}

func (r *reservationRepo) Update(ctx context.Context, reservationDomain domain.Reservation) error {
	reservation := mapper.ReservationDomain2Storage(reservationDomain)
	return r.db.Table("reservations").WithContext(ctx).Save(reservation).Error
}

func (r *reservationRepo) Delete(ctx context.Context, id domain.ReservationID) error {
	return r.db.Table("reservations").WithContext(ctx).Delete(&types.Reservation{}, "id = ?", id).Error
}

func (r *reservationRepo) ListByTourID(ctx context.Context, tourID uuid.UUID) ([]domain.Reservation, error) {
	var reservations []types.Reservation
	err := r.db.Table("reservations").WithContext(ctx).Where("tour_id = ?", tourID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchReservationStorage2Domain(reservations), nil
}

func (r *reservationRepo) ListByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Reservation, error) {
	var reservations []types.Reservation
	err := r.db.Table("reservations").WithContext(ctx).Where("customer_id = ?", userID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchReservationStorage2Domain(reservations), nil
}
