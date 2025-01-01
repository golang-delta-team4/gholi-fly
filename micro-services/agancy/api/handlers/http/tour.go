package http

import (
	pb "gholi-fly-agancy/api/pb"
	"gholi-fly-agancy/api/service"
	"gholi-fly-agancy/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTour(svcGetter ServiceGetter[*service.TourService], cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		// Get AgencyID from path parameter
		agencyID := c.Params("agencyID")
		if _, err := uuid.Parse(agencyID); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid agency ID format")
		}

		var req pb.CreateTourRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Call the service to create a tour
		response, err := svc.CreateTour(
			c.UserContext(),
			agencyID,
			cfg.HotelService.Host,
			cfg.HotelService.Port,
			cfg.TransportService.Host,
			cfg.TransportService.Port,
			&req,
		)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetTour(svcGetter ServiceGetter[*service.TourService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid tour ID format")
		}

		response, err := svc.GetTourByID(c.UserContext(), id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
func UpdateTour(svcGetter ServiceGetter[*service.TourService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid tour ID format")
		}

		var req pb.UpdateTourRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		response, err := svc.UpdateTour(c.UserContext(), id, &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
func DeleteTour(svcGetter ServiceGetter[*service.TourService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid tour ID format")
		}

		response, err := svc.DeleteTour(c.UserContext(), id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
func ListToursByAgency(svcGetter ServiceGetter[*service.TourService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		agencyID := c.Params("agencyID")
		if _, err := uuid.Parse(agencyID); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid agency ID format")
		}

		response, err := svc.ListToursByAgency(c.UserContext(), agencyID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
