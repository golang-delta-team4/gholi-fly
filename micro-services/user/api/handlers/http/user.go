package http

import (
	"errors"
	"user-service/api/handlers/shared"
	"user-service/api/presenter"
	"user-service/api/service"
	"user-service/internal/user"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SignUp(svcGetter shared.ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
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
		ok := presenter.PasswordValidation(req.Password)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, errors.New("invalid password").Error())
		}
		resp, err := svc.SignUp(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, user.ErrEmailNotUnique) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			if errors.Is(err, &service.ErrUserCreationValidation{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(resp)
	}
}

func SignIn(svcGetter shared.ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
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
		accessToken, refreshToken, err := svc.SignIn(c.UserContext(), &req)
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

func Refresh(svcGetter shared.ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userUUID := c.Locals("UserUUID")
		if userUUID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid access token")
		}
		token := userToken(c)
		if token == "" {
			return fiber.NewError(fiber.StatusBadRequest, "refresh token required")
		}
		accessToken, refreshToken, err := svc.Refresh(c.UserContext(), userUUID.(uuid.UUID), token)
		if err != nil {
			if errors.Is(err, service.ErrInvalidRefreshToken) {
				return fiber.NewError(fiber.StatusForbidden, err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(presenter.UserSignInResponse{AccessToken: accessToken, RefreshToken: refreshToken})
	}
}

func GetMe(svcGetter shared.ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userUUID := c.Locals("UserUUID")
		if userUUID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid access token")
		}
		user, err := svc.GetUserByUUID(c.UserContext(), userUUID.(uuid.UUID).String())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(presenter.GetUserResponse{Email: user.Email, FirstName: user.FirstName, LastName: user.LastName})
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
