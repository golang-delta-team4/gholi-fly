package service

import (
	"context"
	"errors"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/hotel"
	"gholi-fly-hotel/internal/hotel/domain"
	hotelPort "gholi-fly-hotel/internal/hotel/port"
)

type HotelService struct {
	svc hotelPort.Service
	// notifSvc              notifPort.Service
}

func NewHotelService(svc hotelPort.Service) *HotelService {
	return &HotelService{
		svc: svc,
	}
}

var (
	ErrHotelCreationValidation = hotel.ErrHotelCreationValidation
	ErrHotelCreationDuplicate  = hotel.ErrHotelCreationDuplicate
	ErrHotelOnCreate           = hotel.ErrHotelCreation
	ErrHotelNotFound           = hotel.ErrHotelNotFound
	ErrInvalidHotelPassword    = errors.New("invalid password")
	ErrWrongOTP                = errors.New("wrong otp")
)

func (s *HotelService) CreateHotel(ctx context.Context, req *pb.HotelCreateRequest) (*pb.HotelCreateResponse, error) {

	hotelId, err := s.svc.CreateHotel(ctx, domain.Hotel{
		Name:       req.Name,
		City:       req.City,
		OwnerEmail: req.OwnerEmail,
	})

	if err != nil {
		return nil, err
	}

	return &pb.HotelCreateResponse{
		HotelId: hotelId.String(),
	}, nil
}
