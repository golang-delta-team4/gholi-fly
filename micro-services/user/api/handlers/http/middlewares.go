package http

import (
	"strconv"
	"user-service/pkg/jwt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			userClaims := userClaims(ctx)
			if userClaims == nil {
				return fiber.ErrUnauthorized
			}
			ctx.Locals("UserID", strconv.Itoa(int(userClaims.UserID)))
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
		AuthScheme: "Bearer",
	})
}