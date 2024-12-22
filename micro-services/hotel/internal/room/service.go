package room

import (
	"context"
	"errors"
	"gholi-fly-hotel/internal/room/domain"
	"gholi-fly-hotel/internal/room/port"
)

var (
	ErrRoomCreation         = errors.New("error on creating room")
	ErrRoomNotFound         = errors.New("room not found")
	ErrInvalidSourceService = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// CreateRoom creates a new room
func (s *service) CreateRoom(ctx context.Context, room domain.Room) (domain.RoomUUID, error) {
	roomID, err := s.repo.Create(ctx, room)
	if err != nil {
		return domain.RoomUUID{}, ErrRoomCreation
	}
	return roomID, nil
}

// GetRoomByID returns a room by its ID
func (s *service) GetRoomByID(ctx context.Context, roomID domain.RoomUUID) (*domain.Room, error) {
	room, err := s.repo.GetByID(ctx, roomID)
	if err != nil {
		return nil, ErrRoomNotFound
	}
	return room, nil
}

// GetRooms returns all rooms
func (s *service) GetRooms(ctx context.Context) ([]domain.Room, error) {
	rooms, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// UpdateRoom updates a room
func (s *service) UpdateRoom(ctx context.Context, room domain.Room) error {
	err := s.repo.Update(ctx, room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room
func (s *service) DeleteRoom(ctx context.Context, roomID domain.RoomUUID) error {
	err := s.repo.Delete(ctx, roomID)
	if err != nil {
		return err
	}
	return nil
}
