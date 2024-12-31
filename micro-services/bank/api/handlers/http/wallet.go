package http

import (
	"errors"
	"gholi-fly-bank/api/pb"
	"gholi-fly-bank/api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetWallets(svcGetter ServiceGetter[*service.WalletService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		// Extract the UserUUID from the context
		userUUID := c.Locals("UserUUID").(uuid.UUID)
		if userUUID == uuid.Nil {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}
		resp, err := svc.GetWallets(c.UserContext(), userUUID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func UpdateWalletBalance(svcGetter ServiceGetter[*service.WalletService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		// Extract the UserUUID from the context
		userUUID := c.Locals("UserUUID").(uuid.UUID)
		if userUUID == uuid.Nil {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		walletID := c.Params("id")
		var req pb.UpdateWalletBalanceRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
		}

		err := svc.UpdateWalletBalance(c.UserContext(), walletID, userUUID, req.DepositeAmount)
		if err != nil {
			if errors.Is(err, service.ErrWalletNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(&pb.UpdateWalletBalanceResponse{
					Status: pb.ResponseStatus_FAILED,
				})
			}
			if errors.Is(err, service.ErrBalanceUpdate) {
				return c.Status(fiber.StatusInternalServerError).JSON(&pb.UpdateWalletBalanceResponse{
					Status: pb.ResponseStatus_FAILED,
				})
			}
		}

		// Return success response
		return c.Status(fiber.StatusOK).JSON(&pb.UpdateWalletBalanceResponse{
			Status: pb.ResponseStatus_SUCCESS,
		})
	}
}
