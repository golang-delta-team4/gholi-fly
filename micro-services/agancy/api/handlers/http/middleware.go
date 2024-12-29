package http

import (
	"gholi-fly-agancy/config"
	"gholi-fly-agancy/pkg/context"
	"gholi-fly-agancy/pkg/logger"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-delta-team4/gholi-fly-shared/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func LoggerMiddleware(cfg config.LoggerConfig) (fiber.Handler, error) {
	logService, err := logger.NewLoggerService(cfg)
	if err != nil {
		return nil, err
	}

	return func(c *fiber.Ctx) error {
		reqID := uuid.New().String()

		// Attach logger for the specific service to the user context
		ctx := logService.AttachLoggerToContext(c.UserContext(), c, reqID)

		// Update Fiber's user context
		c.SetUserContext(ctx)

		// Continue to the next middleware/handler
		return c.Next()
	}, nil
}
