package http

import (
	"errors"
	"user-service/api/presenter"
	"user-service/api/service"
	"user-service/internal/role"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(roleService *service.RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateRoleRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		resp, err := roleService.Create(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, &service.ErrInvalidUUID{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			if errors.Is(err, role.ErrRoleNameNotUnique) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(resp)
	}
}
