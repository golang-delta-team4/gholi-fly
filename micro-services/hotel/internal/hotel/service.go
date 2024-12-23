package hotel

import (
	"context"
	"errors"
	"gholi-fly-hotel/internal/hotel/domain"
	"gholi-fly-hotel/internal/hotel/port"
	"strings"
)

var (
	ErrHotelCreation           = errors.New("error on creating hotel")
	ErrHotelCreationValidation = errors.New("error on creating hotel: validation failed")
	ErrHotelCreationDuplicate  = errors.New("hotel already exists")
	ErrHotelNotFound           = errors.New("hotel not found")
	ErrInvalidSourceService    = errors.New("invalid source service")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// CreateHotel creates a new hotel
func (s *service) CreateHotel(ctx context.Context, hotel domain.Hotel) (domain.HotelUUID, error) {
	hotelID, err := s.repo.Create(ctx, hotel)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return domain.HotelUUID{}, ErrHotelCreationDuplicate
		}
		return domain.HotelUUID{}, ErrHotelCreation
	}
	return hotelID, nil
}

// GetHotelByID returns a hotel by its ID
func (s *service) GetHotelByID(ctx context.Context, hotelID domain.HotelUUID) (*domain.Hotel, error) {
	hotel, err := s.repo.GetByID(ctx, hotelID)
	if err != nil {
		return nil, ErrHotelNotFound
	}
	return hotel, nil
}

// GetHotels returns all hotels
func (s *service) GetHotels(ctx context.Context) ([]domain.Hotel, error) {
	hotels, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

// UpdateHotel updates a hotel
func (s *service) UpdateHotel(ctx context.Context, hotel domain.Hotel) error {
	err := s.repo.Update(ctx, hotel)
	if err != nil {
		return err
	}
	return nil
}

// DeleteHotel deletes a hotel
func (s *service) DeleteHotel(ctx context.Context, hotelID domain.HotelUUID) error {
	err := s.repo.Delete(ctx, hotelID)
	if err != nil {
		return err
	}
	return nil
}
