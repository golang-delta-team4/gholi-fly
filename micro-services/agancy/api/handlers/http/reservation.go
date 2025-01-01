package http

import (
	pb "gholi-fly-agancy/api/pb"
	"gholi-fly-agancy/api/service"
	"gholi-fly-agancy/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateReservation(svcGetter ServiceGetter[*service.ReservationService], cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		// Extract agencyID from path parameters
		agencyID := c.Params("agencyID")
		parsedAgencyID, err := uuid.Parse(agencyID)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid agency ID format")
		}

		// Extract userID from context
		userID, ok := c.Locals("UserUUID").(string)
		if !ok || userID == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "user ID not found in context")
		}
		if _, err := uuid.Parse(userID); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid user ID format")
		}

		// Parse the body into the CreateReservationRequest
		var req pb.CreateReservationRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Validate inputs
		if req.TourId == "" || req.Capacity == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "tourID and capacity are required")
		}
		if _, err := uuid.Parse(req.TourId); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid tour ID format")
		}

		// Call the service to reserve the tour
		reservationID, err := svc.ReserveTour(
			c.UserContext(),
			userID,
			req.TourId,
			uint(req.Capacity),
			parsedAgencyID,
		)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		// Return the created reservation ID
		return c.JSON(fiber.Map{
			"reservation_id": reservationID,
		})
	}
}
