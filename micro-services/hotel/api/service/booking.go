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

func (s *BookingService) CreateBooking(ctx context.Context, req *pb.BookingCreateRequest, hotelId string) (*pb.BookingCreateResponse, error) {
	hotelUUID, err := uuid.Parse(hotelId)
	if err != nil {
		return nil, err
	}

	// Parse all room IDs
	roomUUIDs := make([]uuid.UUID, 0, len(req.RoomIds))
	for _, roomID := range req.RoomIds {
		roomUUID, err := uuid.Parse(roomID)
		if err != nil {
			return nil, ErrBookingCreationValidation
		}
		roomUUIDs = append(roomUUIDs, roomUUID)
	}

	checkIn, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return nil, err
	}
	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return nil, err
	}
	if checkIn.After(checkOut) {
		return nil, ErrBookingCreationValidation
	}

	userUUID, err := uuid.Parse("43ab4a09-b060-4e74-860b-8ab6f1fd1a03")
	if err != nil {
		return nil, ErrBookingCreationValidation
	}
	agencyUUID, err := uuid.Parse("43ab4a09-b060-4e74-860b-9ab6f1fd1a03")
	if err != nil {
		return nil, ErrBookingCreationValidation
	}
	reservationId := uuid.New()
	for _, roomId := range roomUUIDs {
		_, err := s.svc.CreateBookingByHotelID(ctx, domain.Booking{
			CheckIn:       checkIn,
			CheckOut:      checkOut,
			HotelID:       hotelUUID,
			RoomID:        roomId,
			UserID:        &userUUID, // Changed to pointer
			AgencyID:      &agencyUUID,
			ReservationID: reservationId,
			IsPayed:       false,
			Status:        uint8(pb.BookingStatus_BOOKING_PENDING),
		}, hotelUUID)

		if err != nil {
			return nil, err
		}

	}

	return &pb.BookingCreateResponse{
		ReservationId: reservationId.String(),
		TotalPrice:    620000,
	}, nil
}

func (s *BookingService) GetAllBookingsByRoomID(ctx context.Context, roomID string) (*pb.GetAllBookingResponse, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	bookings, err := s.svc.GetAllBookingsByRoomID(ctx, roomUUID)
	if err != nil {
		return nil, err
	}

	var bookingList []*pb.Booking
	for _, r := range bookings {
		booking := &pb.Booking{
			Id:            r.UUID.String(),
			HotelId:       r.HotelID.String(),
			CheckIn:       r.CheckIn.Format("2006-01-02"),
			CheckOut:      r.CheckOut.Format("2006-01-02"),
			BookingStatus: pb.BookingStatus(r.Status),
		}
		if r.UserID != nil {
			booking.UserId = r.UserID.String()
		}
		if r.AgencyID != nil {
			booking.AgencyId = r.AgencyID.String()
		}
		bookingList = append(bookingList, booking)
	}

	return &pb.GetAllBookingResponse{
		Bookings: bookingList,
	}, nil
}

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
