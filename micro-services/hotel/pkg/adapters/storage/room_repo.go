package storage

import (
	"context"
	"errors"

	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	"gholi-fly-hotel/internal/room/domain"
	roomPort "gholi-fly-hotel/internal/room/port"
	"gholi-fly-hotel/pkg/adapters/storage/mapper"
	"gholi-fly-hotel/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepo(db *gorm.DB) roomPort.Repo {
	return &roomRepo{db: db}
}

func (r *roomRepo) CreateByHotelID(ctx context.Context, roomDomain domain.Room, hotelID hotelDomain.HotelUUID) (domain.RoomUUID, error) {
	room := mapper.RoomDomain2Storage(roomDomain)
	room.HotelID = hotelID
	err := r.db.Table("rooms").WithContext(ctx).Create(room).Error
	if err != nil {
		return domain.RoomUUID{}, err
	}
	return domain.RoomUUID(room.UUID), nil
}

func (r *roomRepo) GetByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]domain.Room, error) {
	var rooms []types.Room
	err := r.db.Table("rooms").WithContext(ctx).Where("hotel_id = ?", hotelID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return mapper.BatchRoomStorage2Domain(rooms), nil
}

func (r *roomRepo) GetByID(ctx context.Context, roomID domain.RoomUUID) (*domain.Room, error) {
	var room types.Room

	err := r.db.Table("rooms").WithContext(ctx).Where("uuid = ?", roomID).First(&room).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if room.ID == 0 {
		return nil, nil
	}

	return mapper.RoomStorage2Domain(room), nil
}

func (r *roomRepo) Update(ctx context.Context, roomDomain domain.Room) error {
	room := mapper.RoomDomain2Storage(roomDomain)
	return r.db.Table("rooms").WithContext(ctx).Save(room).Error
}

func (r *roomRepo) Delete(ctx context.Context, roomID domain.RoomUUID) error {
	return r.db.Table("rooms").WithContext(ctx).Delete(&types.Room{}, "uuid = ?", roomID).Error
}
