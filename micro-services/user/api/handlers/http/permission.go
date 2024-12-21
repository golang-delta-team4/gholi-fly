package http

import (
	"user-service/api/presenter"
	"user-service/api/service"

	"github.com/gofiber/fiber/v2"
)

func CreatePermission(permissionService *service.PermissionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreatePermissionRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		resp, err := permissionService.Create(c.UserContext(), &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(resp)
	}
}
