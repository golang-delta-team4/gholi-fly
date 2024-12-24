package storage

import (
	"context"
	"errors"

	"gholi-fly-hotel/internal/hotel/domain"
	hotelPort "gholi-fly-hotel/internal/hotel/port"
	"gholi-fly-hotel/pkg/adapters/storage/mapper"
	"gholi-fly-hotel/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hotelRepo struct {
	db *gorm.DB
}

func NewHotelRepo(db *gorm.DB) hotelPort.Repo {
	return &hotelRepo{db: db}
}

func (r *hotelRepo) Create(ctx context.Context, hotelDomain domain.Hotel) (domain.HotelUUID, error) {
	hotel := mapper.HotelDomain2Storage(hotelDomain)
	err := r.db.Table("hotels").WithContext(ctx).Create(hotel).Error
	if err != nil {
		return domain.HotelUUID{}, err
	}
	return domain.HotelUUID(hotel.UUID), nil
}

func (r *hotelRepo) Get(ctx context.Context) ([]domain.Hotel, error) {
	var hotels []types.Hotel
	err := r.db.Table("hotels").WithContext(ctx).Find(&hotels).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchHotelStorage2Domain(hotels), nil
}

func (r *hotelRepo) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]domain.Hotel, error) {
	var hotels []types.Hotel
	err := r.db.Table("hotels").WithContext(ctx).Where("owner_id = ?", ownerID).Find(&hotels).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchHotelStorage2Domain(hotels), nil
}

func (r *hotelRepo) GetByID(ctx context.Context, hotelID domain.HotelUUID) (*domain.Hotel, error) {
	var hotel types.Hotel

	err := r.db.Table("hotels").WithContext(ctx).Where("uuid = ?", hotelID).First(&hotel).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if hotel.ID == 0 {
		return nil, nil
	}

	return mapper.HotelStorage2Domain(hotel), nil
}

func (r *hotelRepo) Update(ctx context.Context, hotel domain.Hotel) error {
	storageHotel := mapper.HotelDomain2Storage(hotel)
	return r.db.Table("hotels").WithContext(ctx).Save(storageHotel).Error
}

func (r *hotelRepo) Delete(ctx context.Context, hotelID domain.HotelUUID) error {
	return r.db.Table("hotels").WithContext(ctx).Delete(&types.Hotel{}, "uuid = ?", hotelID).Error
}
