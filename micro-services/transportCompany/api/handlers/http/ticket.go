package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/service"
	"github.com/google/uuid"
)

func BuyTicket(svcGetter ServiceGetter[*service.TicketService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userID, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}
		var req pb.BuyTicketRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		response, err := svc.BuyTicket(c.UserContext(), &req, userID)
		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}
