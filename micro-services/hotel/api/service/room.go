package service

import (
	"context"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/room/domain"
	roomPort "gholi-fly-hotel/internal/room/port"

	"github.com/google/uuid"
)

type RoomService struct {
	svc roomPort.Service
	// notifSvc              notifPort.Service
}

func NewRoomService(svc roomPort.Service) *RoomService {
	return &RoomService{
		svc: svc,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *pb.RoomCreateRequest, hotelID string) (*pb.RoomCreateResponse, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	roomId, err := s.svc.CreateRoomByHotelID(ctx, domain.Room{
		HotelID:     hotelUUID,
		RoomNumber:  uint(req.RoomNumber),
		Floor:       uint(req.Floor),
		BasePrice:   uint(req.BasePrice),
		AgencyPrice: uint(req.AgencyPrice),
	}, hotelUUID)

	if err != nil {
		return nil, err
	}

	return &pb.RoomCreateResponse{
		RoomId: roomId.String(),
	}, nil
}
