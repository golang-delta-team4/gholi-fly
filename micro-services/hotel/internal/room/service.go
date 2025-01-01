package room

import (
	"context"
	"errors"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	roomDomain "gholi-fly-hotel/internal/room/domain"
	"gholi-fly-hotel/internal/room/port"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrRoomCreation           = errors.New("error on creating room")
	ErrRoomCreationValidation = errors.New("error on creating room: validation failed")
	ErrRoomCreationDuplicate  = errors.New("room already exists")
	ErrRoomNotFound           = errors.New("room not found")
	ErrInvalidSourceService   = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// CreateRoom creates a new room by hotel id
func (s *service) CreateRoomByHotelID(ctx context.Context, room roomDomain.Room, hotelID hotelDomain.HotelUUID) (roomDomain.RoomUUID, error) {
	if err := room.Validate(); err != nil {
		return uuid.Nil, ErrRoomCreationValidation
	}
	roomID, err := s.repo.CreateByHotelID(ctx, room, hotelID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return roomDomain.RoomUUID{}, ErrRoomCreationDuplicate
		}
		return roomDomain.RoomUUID{}, ErrRoomCreation
	}
	return roomID, nil
}

// GetRooms returns all rooms
func (s *service) GetAllRoomsByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]roomDomain.Room, error) {
	rooms, err := s.repo.GetByHotelID(ctx, hotelID)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetRoomByID returns a room by its ID
func (s *service) GetRoomByID(ctx context.Context, roomID roomDomain.RoomUUID) (*roomDomain.Room, error) {
	room, err := s.repo.GetByID(ctx, roomID)
	if err != nil {
		return nil, ErrRoomNotFound
	}
	return room, nil
}

// UpdateRoom updates a room
func (s *service) UpdateRoom(ctx context.Context, room roomDomain.Room) error {
	err := s.repo.Update(ctx, room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room
func (s *service) DeleteRoom(ctx context.Context, roomID roomDomain.RoomUUID) error {
	err := s.repo.Delete(ctx, roomID)
	if err != nil {
		return err
	}
	return nil
}
