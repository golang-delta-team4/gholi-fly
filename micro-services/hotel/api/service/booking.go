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

func (s *BookingService) CreateUserBooking(ctx context.Context, req *pb.BookingCreateRequest, hotelId string, userUUID domain.UserUUID) (*pb.BookingCreateResponse, error) {
	if req == nil {
		return nil, ErrBookingCreationValidation
	}
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

	reservationId := uuid.New()
	totalPrice := 0
	for _, roomId := range roomUUIDs {
		_, price, err := s.svc.CreateBookingByHotelID(ctx, domain.Booking{
			CheckIn:       checkIn,
			CheckOut:      checkOut,
			HotelID:       hotelUUID,
			RoomID:        roomId,
			UserID:        userUUID,
			ReservationID: reservationId,
			IsPaid:        false,
			Status:        uint8(pb.BookingStatus_BOOKING_PENDING),
		}, hotelUUID, false)

		if err != nil {
			return nil, err
		}
		totalPrice += int(price)

	}

	_, err = s.svc.CreateBookingFactor(ctx, userUUID, hotelUUID, uint(totalPrice), reservationId)
	if err != nil {
		return nil, err
	}

	return &pb.BookingCreateResponse{
		ReservationId: reservationId.String(),
		TotalPrice:    int64(totalPrice),
	}, nil

}

func (s *BookingService) CreateBooking(ctx context.Context, req *pb.BookingCreateRequest, hotelId string, userUUID domain.UserUUID) (*pb.BookingCreateResponse, error) {
	if req == nil {
		return nil, ErrBookingCreationValidation
	}
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

	reservationId := uuid.New()
	totalPrice := 0
	for _, roomId := range roomUUIDs {
		_, price, err := s.svc.CreateBookingByHotelID(ctx, domain.Booking{
			CheckIn:       checkIn,
			CheckOut:      checkOut,
			HotelID:       hotelUUID,
			RoomID:        roomId,
			UserID:        userUUID,
			ReservationID: reservationId,
			IsPaid:        false,
			Status:        uint8(pb.BookingStatus_BOOKING_PENDING),
		}, hotelUUID, true)

		if err != nil {
			return nil, err
		}
		totalPrice += int(price)

	}

	factorId, err := s.svc.CreateBookingFactor(ctx, userUUID, hotelUUID, uint(totalPrice), reservationId)
	if err != nil {
		return nil, err
	}

	return &pb.BookingCreateResponse{
		ReservationId: factorId,
		TotalPrice:    int64(totalPrice),
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
	if bookings == nil || len(bookings) == 0 {
		return nil, ErrBookingNotFound
	}

	var bookingList []*pb.Booking
	for _, r := range bookings {
		booking := &pb.Booking{
			Id:            r.UUID.String(),
			HotelId:       r.HotelID.String(),
			CheckIn:       r.CheckIn.Format("2006-01-02"),
			CheckOut:      r.CheckOut.Format("2006-01-02"),
			BookingStatus: pb.BookingStatus(r.Status),
			UserId:        r.UserID.String(),
			FactorId:      r.FactorID,
		}
		bookingList = append(bookingList, booking)
	}

	return &pb.GetAllBookingResponse{
		Bookings: bookingList,
	}, nil
}

func (s *BookingService) GetAllBookingsByHotelID(ctx context.Context, hotelID string) (*pb.GetAllBookingResponse, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	bookings, err := s.svc.GetAllBookingsByHotelID(ctx, hotelUUID)
	if err != nil {
		return nil, err
	}
	if bookings == nil || len(bookings) == 0 {
		return nil, ErrBookingNotFound
	}

	var bookingList []*pb.Booking
	for _, r := range bookings {
		booking := &pb.Booking{
			Id:            r.UUID.String(),
			HotelId:       r.HotelID.String(),
			CheckIn:       r.CheckIn.Format("2006-01-02"),
			CheckOut:      r.CheckOut.Format("2006-01-02"),
			BookingStatus: pb.BookingStatus(r.Status),
			UserId:        r.UserID.String(),
		}
		bookingList = append(bookingList, booking)
	}

	return &pb.GetAllBookingResponse{
		Bookings: bookingList,
	}, nil
}

func (s *BookingService) GetBookingByID(ctx context.Context, bookingID string) (*pb.Booking, error) {
	bookingUUID, err := uuid.Parse(bookingID)
	if err != nil {
		return nil, err
	}
	booking, err := s.svc.GetBookingByID(ctx, bookingUUID)
	if err != nil {
		return nil, err
	}
	if booking == nil {
		return nil, ErrBookingNotFound
	}

	return &pb.Booking{
		Id:            booking.UUID.String(),
		HotelId:       booking.HotelID.String(),
		CheckIn:       booking.CheckIn.Format("2006-01-02"),
		CheckOut:      booking.CheckOut.Format("2006-01-02"),
		BookingStatus: pb.BookingStatus(booking.Status),
		RoomIds:       []string{booking.RoomID.String()},
		UserId:        booking.UserID.String(),
		FactorId:      booking.FactorID,
	}, nil
}

func (s *BookingService) UpdateBookingStatus(ctx context.Context, bookingID string, status pb.BookingStatus) (*pb.Booking, error) {
	bookingUUID, err := uuid.Parse(bookingID)
	if err != nil {
		return nil, err
	}
	booking, err := s.svc.UpdateBookingStatus(ctx, bookingUUID, uint8(status))
	if err != nil {
		return nil, err
	}

	return &pb.Booking{
		Id:            booking.UUID.String(),
		HotelId:       booking.HotelID.String(),
		CheckIn:       booking.CheckIn.Format("2006-01-02"),
		CheckOut:      booking.CheckOut.Format("2006-01-02"),
		BookingStatus: pb.BookingStatus(booking.Status),
		UserId:        booking.UserID.String(),
	}, nil
}

func (s *BookingService) DeleteBooking(ctx context.Context, bookingID string) error {
	bookingUUID, err := uuid.Parse(bookingID)
	if err != nil {
		return err
	}
	err = s.svc.DeleteBooking(ctx, bookingUUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingService) GetAllBookingsByUserID(ctx context.Context, userUUID uuid.UUID) (*pb.GetAllBookingResponse, error) {

	// userUUID, err := uuid.Parse(userID)
	// if err != nil {
	// 	return nil, err
	// }
	bookings, err := s.svc.GetAllBookingsByUserID(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	if bookings == nil || len(bookings) == 0 {
		return nil, ErrBookingNotFound
	}

	var bookingList []*pb.Booking
	for _, r := range bookings {
		booking := &pb.Booking{
			Id:            r.ReservationID.String(),
			FactorId:      r.FactorID,
			HotelId:       r.HotelID.String(),
			CheckIn:       r.CheckIn.Format("2006-01-02"),
			CheckOut:      r.CheckOut.Format("2006-01-02"),
			BookingStatus: pb.BookingStatus(r.Status),
			UserId:        r.UserID.String(),
		}
		bookingList = append(bookingList, booking)
	}

	return &pb.GetAllBookingResponse{
		Bookings: bookingList,
	}, nil
}

func (s *BookingService) ApproveUserBooking(ctx context.Context, factorID string, userUUID uuid.UUID) error {
	factorUUID, err := uuid.Parse(factorID)
	if err != nil {
		return err
	}
	err = s.svc.ApproveUserBooking(ctx, factorUUID, userUUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingService) CancelUserBooking(ctx context.Context, factorID string, userUUID uuid.UUID) error {
	factorUUID, err := uuid.Parse(factorID)
	if err != nil {
		return err
	}
	err = s.svc.CancelUserBooking(ctx, factorUUID, userUUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingService) CancelBooking(ctx context.Context, factorID string) error {
	factorUUID, err := uuid.Parse(factorID)
	if err != nil {
		return err
	}
	err = s.svc.CancelBooking(ctx, factorUUID)
	if err != nil {
		return err
	}

	return nil
}
