package http

import (
	"errors"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateBookingByRoomID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roomID := c.Params("room_id")
		svc := svcGetter(c.UserContext())
		var req pb.BookingCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.CreateBooking(c.UserContext(), &req, roomID)
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

// func GetAllBookingsByUserID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

// 	}
// }

// func GetBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

// 	}
// }

// func UpdateBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

// 	}
// }

// func DeleteBookingByID(svcGetter ServiceGetter[*service.BookingService]) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

// 	}
// }
