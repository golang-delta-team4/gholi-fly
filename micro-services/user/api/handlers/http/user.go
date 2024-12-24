package http

import (
	"errors"
	"user-service/api/presenter"
	"user-service/api/service"
	"user-service/internal/user"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SignUp(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.UserSignUpRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		err := presenter.EmailValidation(req.Email)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
		var req presenter.UserSignInRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		err := presenter.EmailValidation(req.Email)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		accessToken, refreshToken, err := userService.SignIn(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, user.ErrEmailOrPasswordMismatch) {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}
			if errors.Is(err, user.ErrUserNotFound) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(presenter.UserSignInResponse{AccessToken: accessToken, RefreshToken: refreshToken})
	}
}

func Refresh(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var tokenReq presenter.UserRefreshRequest
		userUUID := c.Locals("UserUUID")
		if userUUID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid access token")
		}
		if err := c.BodyParser(&tokenReq); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "please provide refresh token")
		}
		if tokenReq.RefreshToken == "" {
			return fiber.NewError(fiber.StatusBadRequest, "refresh token required")
		}
		accessToken, err := userService.Refresh(c.UserContext(), userUUID.(uuid.UUID), tokenReq.RefreshToken)
		if err != nil {
			if errors.Is(err, service.ErrInvalidRefreshToken) {
				return fiber.NewError(fiber.StatusForbidden, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(presenter.UserRefreshResponse{AccessToken: accessToken})
	}
}

func validate(req any) map[string]string {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, validationError := range errs {
			return map[string]string{
				"field":   validationError.Field(),
				"message": validationError.Error(),
			}
		}
	}
	return nil
}
