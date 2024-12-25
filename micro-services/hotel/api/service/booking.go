package service

import (
	"context"
	"time"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/booking"
	"gholi-fly-hotel/internal/booking/domain"
	bookingPort "gholi-fly-hotel/internal/booking/port"

	"github.com/google/uuid"
)

type BookingService struct {
	svc bookingPort.Service
	// notifSvc              notifPort.Service
}

func NewBookingService(svc bookingPort.Service) *BookingService {
	return &BookingService{
		svc: svc,
	}
}

var (
	ErrBookingCreationValidation = booking.ErrBookingCreationValidation
	ErrBookingCreationDuplicate  = booking.ErrBookingCreationDuplicate
	ErrBookingOnCreate           = booking.ErrBookingCreation
	ErrBookingNotFound           = booking.ErrBookingNotFound
)

func (s *BookingService) CreateBooking(ctx context.Context, req *pb.BookingCreateRequest, roomId string) (*pb.BookingCreateResponse, error) {
	roomUUID, err := uuid.Parse(roomId)
	if err != nil {
		return nil, err
	}
	hotelUUID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return nil, err
	}
	checkIn, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return nil, err
	}
	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return nil, err
	}
	bookingId, err := s.svc.CreateBookingByRoomID(ctx, domain.Booking{
		CheckIn:  checkIn,
		CheckOut: checkOut,
		HotelID:  hotelUUID,
		Status:   1,
	}, roomUUID)

	if err != nil {
		return nil, err
	}

	return &pb.BookingCreateResponse{
		BookingId: bookingId.String(),
	}, nil
}

// func (s *RoomService) GetAllRooms(ctx context.Context, hotelID string) (*pb.GetAllRoomsResponse, error) {
// 	hotelUUID, err := uuid.Parse(hotelID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rooms, err := s.svc.GetAllRoomsByHotelID(ctx, hotelUUID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var roomList []*pb.Room
// 	for _, r := range rooms {
// 		roomList = append(roomList, &pb.Room{
// 			Id:          r.UUID.String(),
// 			HotelId:     r.HotelID.String(),
// 			RoomNumber:  int32(r.RoomNumber),
// 			Floor:       int32(r.Floor),
// 			BasePrice:   int32(r.BasePrice),
// 			AgencyPrice: int32(r.AgencyPrice),
// 		})
// 	}

// 	return &pb.GetAllRoomsResponse{
// 		Rooms: roomList,
// 	}, nil
// }

// func (s *RoomService) GetRoomByID(ctx context.Context, roomID string) (*pb.Room, error) {
// 	roomUUID, err := uuid.Parse(roomID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	room, err := s.svc.GetRoomByID(ctx, roomUUID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Room{
// 		Id:          room.UUID.String(),
// 		HotelId:     room.HotelID.String(),
// 		RoomNumber:  int32(room.RoomNumber),
// 		Floor:       int32(room.Floor),
// 		BasePrice:   int32(room.BasePrice),
// 		AgencyPrice: int32(room.AgencyPrice),
// 	}, nil
// }
