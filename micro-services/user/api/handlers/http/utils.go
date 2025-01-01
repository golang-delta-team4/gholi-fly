package http

import (
	"github.com/gofiber/fiber/v2"
	goJwt "github.com/golang-jwt/jwt/v5"
)

func userToken(ctx *fiber.Ctx) string {
	if u := ctx.Locals("user"); u != nil {
		return u.(*goJwt.Token).Raw
	}
	return ""
}
