package booking

import (
	"context"
	"errors"
	bookingDomain "gholi-fly-hotel/internal/booking/domain"
	"gholi-fly-hotel/internal/booking/port"
	hotelDomain "gholi-fly-hotel/internal/hotel/domain"
	hotelPort "gholi-fly-hotel/internal/hotel/port"
	roomDomain "gholi-fly-hotel/internal/room/domain"
	bankPb "gholi-fly-hotel/pkg/adapters/clients/grpc/pb"
	bankClientPort "gholi-fly-hotel/pkg/adapters/clients/grpc/port"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrBookingCreation           = errors.New("error on creating booking")
	ErrBookingCreationValidation = errors.New("error on creating booking: validation failed")
	ErrBookingCreationDuplicate  = errors.New("booking already exists in these days")
	ErrBookingNotFound           = errors.New("booking not found")
	ErrInvalidSourceService      = errors.New("invalid source service")
	ErrBookingApprovalFailed     = errors.New("error on approving booking")
	ErrBookingCancellationFailed = errors.New("error on cancelling booking")
)

type service struct {
	repo        port.Repo
	hotelRepo   hotelPort.Repo
	bankClient  bankClientPort.GRPCBankClient
	notifClient bankClientPort.GRPCNotificationClient
}

func NewService(repo port.Repo, hotelRepo hotelPort.Repo, bankClient bankClientPort.GRPCBankClient, notifClient bankClientPort.GRPCNotificationClient) port.Service {
	return &service{
		repo:        repo,
		hotelRepo:   hotelRepo,
		bankClient:  bankClient,
		notifClient: notifClient,
	}
}

// CreateBookingByRoomID creates a new booking by room ID
func (s *service) CreateBookingByHotelID(ctx context.Context, booking bookingDomain.Booking, hotelID hotelDomain.HotelUUID, isAgency bool) (bookingDomain.BookingUUID, roomDomain.RoomPrice, error) {
	if err := booking.Validate(); err != nil {
		return uuid.Nil, 0, ErrBookingCreationValidation
	}
	bookingID, price, err := s.repo.CreateByHotelID(ctx, booking, hotelID, isAgency)
	if err != nil {
		if strings.Contains(err.Error(), ErrBookingCreationDuplicate.Error()) {
			return bookingDomain.BookingUUID{}, 0, ErrBookingCreationDuplicate
		}
		return bookingDomain.BookingUUID{}, 0, ErrBookingCreation
	}

	return bookingID, price, nil
}

func (s *service) CreateBookingFactor(ctx context.Context, userId uuid.UUID, hotelID hotelDomain.HotelUUID, totalPrice uint, bookingId bookingDomain.BookingUUID) (string, error) {

	hotel, err := s.hotelRepo.GetByID(ctx, hotelID)
	if err != nil {
		return "", err
	}
	ownerId := hotel.OwnerID

	walletResponse, err := s.bankClient.GetUserWallets(&bankPb.GetWalletsRequest{
		OwnerId: ownerId.String(),
	})
	if walletResponse == nil || err != nil {
		return "", ErrBookingCreation
	}

	response, err := s.bankClient.CreateFactor(&bankPb.CreateFactorRequest{
		Factor: &bankPb.Factor{
			SourceService: "Hotel_Service",
			TotalAmount:   uint64(totalPrice),
			CustomerId:    userId.String(),
			BookingId:     bookingId.String(),
			ExternalId:    bookingId.String(),

			Distributions: []*bankPb.Distribution{
				{
					WalletId: walletResponse.Wallets[0].Id,
					Amount:   uint64(totalPrice),
				},
			},
		},
	})
	if err != nil {
		return "", err
	}
	s.repo.AddBookingFactor(ctx, bookingId, response.Factor.Id)

	// notifResponse, err := s.notifClient.AddNotification(&bankPb.AddNotificationRequest{
	// 	EventName: "BookingCreated",
	// 	UserId:    userId.String(),
	// 	Message:   "Booking created successfully",
	// })
	// if notifResponse == nil || err != nil {
	// 	return "", ErrBookingCreation
	// }

	return response.Factor.Id, nil
}

func (s *service) ApproveUserBooking(ctx context.Context, factorID uuid.UUID, userUUID uuid.UUID) error {

	resp, err := s.bankClient.ApplyFactor(&bankPb.ApplyFactorRequest{
		FactorId: factorID.String(),
	})
	if err != nil || resp.Status != bankPb.ResponseStatus_SUCCESS {
		return ErrBookingApprovalFailed
	}

	err = s.repo.ApproveUserBooking(ctx, factorID, userUUID)
	if err != nil {
		return ErrBookingApprovalFailed
	}

	// notifResponse, err := s.notifClient.AddNotification(&bankPb.AddNotificationRequest{
	// 	EventName: "BookingApprove",
	// 	UserId:    userUUID.String(),
	// 	Message:   "Booking payed successfully",
	// })
	// if notifResponse == nil || err != nil {
	// 	return ErrBookingCreation
	// }

	return nil
}

func (s *service) CancelUserBooking(ctx context.Context, factorID uuid.UUID, userUUID uuid.UUID) error {
	resp, err := s.bankClient.CancelFactor(&bankPb.CancelFactorRequest{
		FactorId: factorID.String(),
	})
	if err != nil || resp.Status != bankPb.ResponseStatus_SUCCESS {
		return ErrBookingApprovalFailed
	}
	err = s.repo.CancelUserBooking(ctx, factorID, userUUID)
	if err != nil {
		return ErrBookingCancellationFailed
	}

	// notifResponse, err := s.notifClient.AddNotification(&bankPb.AddNotificationRequest{
	// 	EventName: "BookingCancel",
	// 	UserId:    userUUID.String(),
	// 	Message:   "Booking cancelled",
	// })
	// if notifResponse == nil || err != nil {
	// 	return ErrBookingCreation
	// }
	return nil
}

func (s *service) CancelBooking(ctx context.Context, factorID uuid.UUID) error {
	resp, err := s.bankClient.CancelFactor(&bankPb.CancelFactorRequest{
		FactorId: factorID.String(),
	})
	if err != nil || resp.Status != bankPb.ResponseStatus_SUCCESS {
		return ErrBookingApprovalFailed
	}
	err = s.repo.CancelBooking(ctx, factorID)
	if err != nil {
		return ErrBookingCancellationFailed
	}
	return nil
}

// GetAllBookingsByRoomID returns all bookings by room ID
func (s *service) GetAllBookingsByRoomID(ctx context.Context, roomID roomDomain.RoomUUID) ([]bookingDomain.Booking, error) {
	return s.repo.GetByRoomID(ctx, roomID)
}

// GetAllBookingsByUserID returns all bookings by user ID
func (s *service) GetAllBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]bookingDomain.Booking, error) {
	return s.repo.GetByUserID(ctx, userID)
}

// GetAllBookingsByHotelID returns all bookings by hotel ID
func (s *service) GetAllBookingsByHotelID(ctx context.Context, hotelID hotelDomain.HotelUUID) ([]bookingDomain.Booking, error) {
	return s.repo.GetAllBookingsByHotelID(ctx, hotelID)
}

// GetBookingByID returns a booking by its ID
func (s *service) GetBookingByID(ctx context.Context, bookingID bookingDomain.BookingUUID) (*bookingDomain.Booking, error) {
	booking, err := s.repo.GetByID(ctx, bookingID)
	if err != nil {
		return nil, ErrBookingNotFound
	}
	return booking, nil
}

// UpdateBooking updates a booking
func (s *service) UpdateBooking(ctx context.Context, booking bookingDomain.Booking) error {
	return s.repo.Update(ctx, booking)
}

// UpdateBookingStatus updates the status of a booking
func (s *service) UpdateBookingStatus(ctx context.Context, bookingID bookingDomain.BookingUUID, status uint8) (*bookingDomain.Booking, error) {
	booking, err := s.repo.GetByID(ctx, bookingID)
	if err != nil {
		return nil, ErrBookingNotFound
	}
	booking.Status = status
	err = s.repo.Update(ctx, *booking)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

// DeleteBooking deletes a booking
func (s *service) DeleteBooking(ctx context.Context, bookingID bookingDomain.BookingUUID) error {
	return s.repo.Delete(ctx, bookingID)
}
