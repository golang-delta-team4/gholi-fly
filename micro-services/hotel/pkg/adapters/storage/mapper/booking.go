package mapper

import (
	"gholi-fly-hotel/internal/booking/domain"
	"gholi-fly-hotel/pkg/adapters/storage/types"
	"gholi-fly-hotel/pkg/fp"

	"gorm.io/gorm"
)

func BookingDomain2Storage(bookingDomain domain.Booking) *types.Booking {
	return &types.Booking{
		Model: gorm.Model{
			ID:        uint(bookingDomain.ID),
			CreatedAt: bookingDomain.CreatedAt,
			UpdatedAt: bookingDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(bookingDomain.DeletedAt)),
		},
		UUID:    bookingDomain.UUID,
		HotelID: bookingDomain.HotelID,
		RoomID:  bookingDomain.RoomID,
		UserID:  bookingDomain.UserID,
		// AgencyID:      bookingDomain.AgencyID,
		ReservationID: bookingDomain.ReservationID,
		IsPayed:       bookingDomain.IsPayed,
		CheckIn:       bookingDomain.CheckIn,
		CheckOut:      bookingDomain.CheckOut,
		Status:        bookingDomain.Status,
	}
}

func bookingDomain2Storage(bookingDomain domain.Booking) types.Booking {
	return types.Booking{
		Model: gorm.Model{
			ID:        uint(bookingDomain.ID),
			CreatedAt: bookingDomain.CreatedAt,
			UpdatedAt: bookingDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(bookingDomain.DeletedAt)),
		},
		UUID:    bookingDomain.UUID,
		HotelID: bookingDomain.HotelID,
		RoomID:  bookingDomain.RoomID,
		UserID:  bookingDomain.UserID,
		// AgencyID:      bookingDomain.AgencyID,
		ReservationID: bookingDomain.ReservationID,
		IsPayed:       bookingDomain.IsPayed,
		CheckIn:       bookingDomain.CheckIn,
		CheckOut:      bookingDomain.CheckOut,
		Status:        bookingDomain.Status,
	}
}

func BatchBookingDomain2Storage(domains []domain.Booking) []types.Booking {
	return fp.Map(domains, bookingDomain2Storage)
}

func BookingStorage2Domain(booking types.Booking) *domain.Booking {
	return &domain.Booking{
		ID:            domain.BookingID(booking.ID),
		UUID:          domain.BookingUUID(booking.UUID),
		HotelID:       booking.HotelID,
		RoomID:        booking.RoomID,
		UserID:        booking.UserID,
		FactorID:      booking.FactorID,
		ReservationID: booking.ReservationID,
		IsPayed:       booking.IsPayed,
		CheckIn:       booking.CheckIn,
		CheckOut:      booking.CheckOut,
		Status:        booking.Status,
		CreatedAt:     booking.CreatedAt,
		UpdatedAt:     booking.UpdatedAt,
		DeletedAt:     booking.DeletedAt.Time,
	}
}

func bookingStorage2Domain(booking types.Booking) domain.Booking {
	return domain.Booking{
		ID:            domain.BookingID(booking.ID),
		UUID:          domain.BookingUUID(booking.UUID),
		HotelID:       booking.HotelID,
		RoomID:        booking.RoomID,
		UserID:        booking.UserID,
		FactorID:      booking.FactorID,
		ReservationID: booking.ReservationID,
		IsPayed:       booking.IsPayed,
		CheckIn:       booking.CheckIn,
		CheckOut:      booking.CheckOut,
		Status:        booking.Status,
		CreatedAt:     booking.CreatedAt,
		UpdatedAt:     booking.UpdatedAt,
		DeletedAt:     booking.DeletedAt.Time,
	}
}

func BatchBookingStorage2Domain(bookings []types.Booking) []domain.Booking {
	return fp.Map(bookings, bookingStorage2Domain)
}
