package http

import (

	"github.com/golang-delta-team4/gholi-fly-shared/jwt"

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