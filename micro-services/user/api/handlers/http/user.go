package http

import (
	"errors"
	"user-service/api/service"
	"user-service/internal/user"
	"user-service/internal/user/domain"

	"github.com/gofiber/fiber/v2"
)

func SignUp(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req domain.User
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		resp, err := userService.SignUp(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, &service.ErrUserCreationValidation{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(resp)
	}
}

func SignIn(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req domain.UserSignInRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		accessToken, refreshToken, err := userService.SignIn(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, user.ErrEmailOrPasswordMismatch{}) {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}
			if errors.Is(err, user.ErrUserNotFound{}) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(fiber.Map{"refreshToken": refreshToken, "accessToken": accessToken})
	}
}
