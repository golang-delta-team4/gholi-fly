package service

import (
	"context"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/hotel"
	"gholi-fly-hotel/internal/hotel/domain"
	hotelPort "gholi-fly-hotel/internal/hotel/port"

	"github.com/google/uuid"
)

type HotelService struct {
	svc hotelPort.Service
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
)

func (s *HotelService) CreateHotel(ctx context.Context, req *pb.HotelCreateRequest) (*pb.HotelCreateResponse, error) {
	ownerUUID, err := uuid.Parse(req.OwnerId)
	if err != nil {
		return nil, err
	}
	hotelId, err := s.svc.CreateHotel(ctx, domain.Hotel{
		Name:    req.Name,
		City:    req.City,
		OwnerID: ownerUUID,
	})

	if err != nil {
		return nil, err
	}

	return &pb.HotelCreateResponse{
		HotelId: hotelId.String(),
	}, nil
}

func (s *HotelService) GetAllHotels(ctx context.Context) (*pb.GetAllHotelsResponse, error) {
	hotels, err := s.svc.GetAllHotels(ctx)
	if err != nil {
		return nil, err
	}

	var hotelList []*pb.Hotel
	for _, h := range hotels {
		var rooms []*pb.Room
		for _, r := range h.Rooms {
			rooms = append(rooms, &pb.Room{
				Id:          r.UUID.String(),
				RoomNumber:  int32(r.RoomNumber),
				Floor:       int32(r.Floor),
				BasePrice:   int32(r.BasePrice),
				AgencyPrice: int32(r.AgencyPrice),
			})
		}
		hotelList = append(hotelList, &pb.Hotel{
			Id:    h.UUID.String(),
			Name:  h.Name,
			City:  h.City,
			Rooms: rooms,
		})
	}

	return &pb.GetAllHotelsResponse{
		Hotels: hotelList,
	}, nil
}

func (s *HotelService) GetAllHotelsByOwnerID(ctx context.Context, ownerID string) (*pb.GetAllHotelsResponse, error) {
	ownerUUID, err := uuid.Parse(ownerID)
	if err != nil {
		return nil, err
	}
	hotels, err := s.svc.GetAllHotelsByOwnerID(ctx, ownerUUID)
	if err != nil {
		return nil, err
	}

	var hotelList []*pb.Hotel
	for _, h := range hotels {
		var rooms []*pb.Room
		for _, r := range h.Rooms {
			rooms = append(rooms, &pb.Room{
				Id:          r.UUID.String(),
				RoomNumber:  int32(r.RoomNumber),
				Floor:       int32(r.Floor),
				BasePrice:   int32(r.BasePrice),
				AgencyPrice: int32(r.AgencyPrice),
			})
		}
		hotelList = append(hotelList, &pb.Hotel{
			Id:    h.UUID.String(),
			Name:  h.Name,
			City:  h.City,
			Rooms: rooms,
		})
	}

	return &pb.GetAllHotelsResponse{
		Hotels: hotelList,
	}, nil
}

func (s *HotelService) GetHotelByID(ctx context.Context, hotelID string) (*pb.Hotel, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	hotel, err := s.svc.GetHotelByID(ctx, hotelUUID)
	if err != nil {
		return nil, err
	}
	var rooms []*pb.Room
	for _, r := range hotel.Rooms {
		rooms = append(rooms, &pb.Room{
			Id:          r.UUID.String(),
			RoomNumber:  int32(r.RoomNumber),
			Floor:       int32(r.Floor),
			BasePrice:   int32(r.BasePrice),
			AgencyPrice: int32(r.AgencyPrice),
		})
	}
	return &pb.Hotel{
		Id:    hotel.UUID.String(),
		Name:  hotel.Name,
		City:  hotel.City,
		Rooms: rooms,
	}, nil
}

func (s *HotelService) UpdateHotel(ctx context.Context, req *pb.UpdateHotelRequest, hotelID string) (*pb.UpdateHotelResponse, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	err = s.svc.UpdateHotel(ctx, domain.Hotel{
		UUID: domain.HotelUUID(hotelUUID),
		Name: req.Name,
		City: req.City,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateHotelResponse{}, nil
}

func (s *HotelService) DeleteHotel(ctx context.Context, hotelID string) (*pb.DeleteHotelResponse, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	err = s.svc.DeleteHotel(ctx, domain.HotelUUID(hotelUUID))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteHotelResponse{}, nil
}
