package service

import (
	"context"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/room"
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

var (
	ErrRoomCreationValidation = room.ErrRoomCreationValidation
	ErrRoomCreationDuplicate  = room.ErrRoomCreationDuplicate
	ErrRoomOnCreate           = room.ErrRoomCreation
	ErrRoomNotFound           = room.ErrRoomNotFound
)

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

func (s *RoomService) GetAllRooms(ctx context.Context, hotelID string) (*pb.GetAllRoomsResponse, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	rooms, err := s.svc.GetAllRoomsByHotelID(ctx, hotelUUID)
	if err != nil {
		return nil, err
	}

	var roomList []*pb.Room
	for _, r := range rooms {
		roomList = append(roomList, &pb.Room{
			Id:          r.UUID.String(),
			HotelId:     r.HotelID.String(),
			RoomNumber:  int32(r.RoomNumber),
			Floor:       int32(r.Floor),
			BasePrice:   int32(r.BasePrice),
			AgencyPrice: int32(r.AgencyPrice),
		})
	}

	return &pb.GetAllRoomsResponse{
		Rooms: roomList,
	}, nil
}

func (s *RoomService) GetRoomByID(ctx context.Context, roomID string) (*pb.Room, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	room, err := s.svc.GetRoomByID(ctx, roomUUID)
	if err != nil {
		return nil, err
	}

	return &pb.Room{
		Id:          room.UUID.String(),
		HotelId:     room.HotelID.String(),
		RoomNumber:  int32(room.RoomNumber),
		Floor:       int32(room.Floor),
		BasePrice:   int32(room.BasePrice),
		AgencyPrice: int32(room.AgencyPrice),
	}, nil
}
