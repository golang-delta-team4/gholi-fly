package http

import (
	"user-service/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	goJwt "github.com/golang-jwt/jwt/v5"
)

func userClaims(ctx *fiber.Ctx) *jwt.UserClaims {
	if u := ctx.Locals("user"); u != nil {
		userClaims, ok := u.(*goJwt.Token).Claims.(*jwt.UserClaims)
		if ok {
			return userClaims
		}
	}
	return nil
}
