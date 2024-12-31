package http

import (
	"strings"
	"user-service/api/presenter"
	"user-service/api/service"
	"user-service/pkg/context"
	"user-service/pkg/logger"

	"gorm.io/gorm"

	"github.com/golang-delta-team4/gholi-fly-shared/jwt"
	"github.com/google/uuid"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			userClaims := jwt.GetUserClaims(ctx)
			if userClaims == nil {
				return fiber.ErrUnauthorized
			}
			ctx.Locals("UserUUID", userClaims.UserUUID)
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
		AuthScheme: "Bearer",
	})
}

func setUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext(), context.WithLogger(logger.NewLogger())))
	return c.Next()
}

func setTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tx := db.Begin()

		context.SetDB(c.UserContext(), tx, true)

		if err := c.Next(); err != nil {
			context.Rollback(c.UserContext())
			return err
		}
		if c.Response().StatusCode() >= 300 {
			return context.Rollback(c.UserContext())
		}

		if err := context.CommitOrRollback(c.UserContext(), true); err != nil {
			return err
		}

		return nil
	}
}

func newAuthorizationMiddlewareDirect(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userUUID := c.Locals("UserUUID")
		if userUUID == nil {
			return fiber.ErrUnauthorized
		}
		routeDetail := strings.Split(c.Path(), "/")
		lastPart := strings.Split(routeDetail[len(routeDetail)-1], "?")
		ok, err := UserAuthorization(userService, c.UserContext(), presenter.UserAuthorization{
			UserUUID: userUUID.(uuid.UUID),
			Route:    strings.Join(routeDetail[3:len(routeDetail)-1], "/")+"/"+lastPart[0],
			Method:   c.Method()})
		if err != nil {
			return err
		}
		if ok {
			return c.Next()
		}
		return fiber.ErrForbidden
	}
}
