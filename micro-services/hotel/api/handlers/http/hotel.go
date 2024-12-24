package http

import (
	"errors"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateHotel(svcGetter ServiceGetter[*service.HotelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.HotelCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.CreateHotel(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrHotelCreationValidation) || errors.Is(err, service.ErrHotelCreationDuplicate) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetAllHotels(svcGetter ServiceGetter[*service.HotelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		resp, err := svc.GetAllHotels(c.UserContext())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetAllHotelsByOwnerID(svcGetter ServiceGetter[*service.HotelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		ownerID := c.Params("owner_id")

		resp, err := svc.GetAllHotelsByOwnerID(c.UserContext(), ownerID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetHotelByID(svcGetter ServiceGetter[*service.HotelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		hotelID := c.Params("id")

		resp, err := svc.GetHotelByID(c.UserContext(), hotelID)
		if err != nil {
			if errors.Is(err, service.ErrHotelNotFound) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func UpdateHotelByID(svcGetter ServiceGetter[*service.HotelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		hotelID := c.Params("id")
		var req pb.UpdateHotelRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.UpdateHotel(c.UserContext(), &req, hotelID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func DeleteHotelByID(svcGetter ServiceGetter[*service.HotelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		hotelID := c.Params("id")

		resp, err := svc.DeleteHotel(c.UserContext(), hotelID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
