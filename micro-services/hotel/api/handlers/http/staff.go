package http

import (
	"errors"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateStaffByHotelID(svcGetter ServiceGetter[*service.StaffService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hotelID := c.Params("hotel_id")
		svc := svcGetter(c.UserContext())
		var req pb.StaffCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.CreateStaff(c.UserContext(), &req, hotelID)
		if err != nil {
			if errors.Is(err, service.ErrHotelCreationValidation) || errors.Is(err, service.ErrHotelCreationDuplicate) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)

	}
}

func GetAllStaffsByHotelID(svcGetter ServiceGetter[*service.StaffService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

	}
}

func GetStaffByID(svcGetter ServiceGetter[*service.StaffService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

	}
}

func UpdateStaffByID(svcGetter ServiceGetter[*service.StaffService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

	}
}

func DeleteStaffByID(svcGetter ServiceGetter[*service.StaffService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "err.Error()")

	}
}
