package http

import (
	"errors"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUserBookingByHotelID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hotelID := c.Params("hotel_id")
		svc := svcGetter(c.UserContext())
		var req pb.BookingCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		userID, ok := c.Locals("UserUUID").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}
		resp, err := svc.CreateUserBooking(c.UserContext(), &req, hotelID, userID)
		if err != nil {
			if errors.Is(err, service.ErrBookingCreationValidation) || errors.Is(err, service.ErrBookingCreationDuplicate) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)

	}
}

func CreateBookingByHotelID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hotelID := c.Params("hotel_id")
		svc := svcGetter(c.UserContext())
		var req pb.BookingCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		// hotelID := req.HotelId
		userId := req.UserId
		userUUID, err := uuid.Parse(userId)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.CreateBooking(c.UserContext(), &req, hotelID, userUUID)
		if err != nil {
			if errors.Is(err, service.ErrBookingCreationValidation) || errors.Is(err, service.ErrBookingCreationDuplicate) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)

	}
}

func GetAllBookingsByRoomID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		roomId := c.Params("room_id")

		resp, err := svc.GetAllBookingsByRoomID(c.UserContext(), roomId)
		if err != nil {
			if errors.Is(err, service.ErrBookingNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetAllBookingsByHotelID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		hotelID := c.Params("hotel_id")

		resp, err := svc.GetAllBookingsByHotelID(c.UserContext(), hotelID)
		if err != nil {
			if errors.Is(err, service.ErrBookingNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetAllBookingsByUserID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userID, ok := c.Locals("UserUUID").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		resp, err := svc.GetAllBookingsByUserID(c.UserContext(), userID)
		if err != nil {
			if errors.Is(err, service.ErrBookingNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		bookingID := c.Params("id")

		resp, err := svc.GetBookingByID(c.UserContext(), bookingID)
		if err != nil {
			if errors.Is(err, service.ErrBookingNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func CancelUserBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		bookingID := c.Params("factor_id")
		userID, ok := c.Locals("UserUUID").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		err := svc.CancelUserBooking(c.UserContext(), bookingID, userID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func CancelBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		bookingID := c.Params("factor_id")

		err := svc.CancelBooking(c.UserContext(), bookingID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func ApproveUserBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		bookingID := c.Params("factor_id")
		userID, ok := c.Locals("UserUUID").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		err := svc.ApproveUserBooking(c.UserContext(), bookingID, userID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func DeleteBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		bookingID := c.Params("booking_id")

		err := svc.DeleteBooking(c.UserContext(), bookingID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
