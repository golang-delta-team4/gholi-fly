package storage

import (
	"context"
	"errors"

	"gholi-fly-hotel/internal/booking/domain"
	bookingPort "gholi-fly-hotel/internal/booking/port"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"
	"gholi-fly-hotel/pkg/adapters/storage/mapper"
	"gholi-fly-hotel/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookingRepo struct {
	db *gorm.DB
}

func NewBookingRepo(db *gorm.DB) bookingPort.Repo {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) CreateByHotelID(ctx context.Context, bookingDomain domain.Booking, hotelID hotelDomain.HotelUUID) (domain.BookingUUID, error) {
	booking := mapper.BookingDomain2Storage(bookingDomain)
	booking.HotelID = hotelID
	booking.RoomIDs = bookingDomain.RoomIDs

	err := r.db.Table("bookings").WithContext(ctx).Create(booking).Error
	if err != nil {
		return domain.BookingUUID{}, err
	}

	return domain.BookingUUID(booking.UUID), nil
}

func (r *bookingRepo) GetByID(ctx context.Context, bookingID domain.BookingUUID) (*domain.Booking, error) {
	var booking types.Booking

	err := r.db.Table("bookings").
		Preload("Hotel").
		WithContext(ctx).
		Where("uuid = ?", bookingID).
		First(&booking).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if booking.ID == 0 {
		return nil, nil
	}

	return mapper.BookingStorage2Domain(booking), nil
}

func (r *bookingRepo) GetByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]domain.Booking, error) {
	var bookings []types.Booking
	err := r.db.Table("bookings").WithContext(ctx).Where("room_id = ?", roomID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchBookingStorage2Domain(bookings), nil
}

func (r *bookingRepo) GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Booking, error) {
	var bookings []types.Booking
	err := r.db.Table("bookings").WithContext(ctx).Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchBookingStorage2Domain(bookings), nil
}

func (r *bookingRepo) Update(ctx context.Context, bookingDomain domain.Booking) error {
	booking := mapper.BookingDomain2Storage(bookingDomain)
	return r.db.Table("bookings").WithContext(ctx).Save(booking).Error
}

func (r *bookingRepo) Delete(ctx context.Context, bookingID domain.BookingUUID) error {
	return r.db.Table("bookings").WithContext(ctx).Delete(&types.Booking{}, "uuid = ?", bookingID).Error
}
