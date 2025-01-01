package http

import (
	pb "gholi-fly-agancy/api/pb"
	"gholi-fly-agancy/api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userID, ok := c.Locals("UserUUID").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		var req pb.CreateAgencyRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		response, err := svc.CreateAgency(c.UserContext(), userID.String(), &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
func GetAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("im here 🙋🏻‍♂️")
		svc := svcGetter(c.UserContext())

		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid agency ID format")
		}

		response, err := svc.GetAgency(c.UserContext(), id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
func UpdateAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid agency ID format")
		}

		userID, ok := c.Locals("UserUUID").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		var req pb.UpdateAgencyRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		response, err := svc.UpdateAgency(c.UserContext(), id, userID.String(), &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
func DeleteAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		// Get the agency ID from the path parameters
		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid agency ID format")
		}

		// Call the service to delete the agency
		_, err := svc.DeleteAgency(c.UserContext(), id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
