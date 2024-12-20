package http

import (
	"errors"
	"strconv"
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

func Refresh(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type tokenReqBody struct {
			RefreshToken string `json:"refreshToken"`
		}
		var tokenReq tokenReqBody
		userID, err := strconv.Atoi(c.Locals("UserID").(string))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "token invalid")
		}
		if err := c.BodyParser(&tokenReq); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "please provide refresh token")
		}
		if tokenReq.RefreshToken == "" {
			return fiber.NewError(fiber.StatusBadRequest, "refresh token required")
		}
		accessToken, err := userService.Refresh(c.UserContext(), uint(userID), tokenReq.RefreshToken)
		if err != nil {
			if errors.Is(err, service.ErrInvalidRefreshToken) {
				return fiber.NewError(fiber.StatusForbidden, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(fiber.Map{"accessToken": accessToken})
	}
}
