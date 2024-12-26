package http

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/jwt"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/logger"
	"github.com/google/uuid"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/context"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//func newAuthMiddleware(secret []byte) fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:  jwtware.SigningKey{Key: secret},
// 		Claims:      &jwt.UserClaims{},
// 		TokenLookup: "header:Authorization",
// 		SuccessHandler: func(ctx *fiber.Ctx) error {
// 			userClaims := userClaims(ctx)
// 			if userClaims == nil {
// 				return fiber.ErrUnauthorized
// 			}

// 			logger := context.GetLogger(ctx.UserContext())
// 			context.SetLogger(ctx.UserContext(), logger.With("user_id", userClaims.UserID))

// 			return ctx.Next()
// 		},
// 		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
// 			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
// 		},
// 		AuthScheme: "Bearer",
// 	})
// }

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			ctx.Locals("user_id", uuid.MustParse("f28e4b35-00e8-46e7-817f-b4f5908f3146"))
			return ctx.Next()
			userClaims := userClaims(ctx)
			if userClaims == nil {
				return fiber.ErrUnauthorized
			}

			logger := context.GetLogger(ctx.UserContext())
			context.SetLogger(ctx.UserContext(), logger.With("user_id", userClaims.UserID))

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
