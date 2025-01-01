package jwt

import (
	"github.com/gofiber/fiber/v2"
	goJwt "github.com/golang-jwt/jwt/v5"
)

func GetUserClaims(ctx *fiber.Ctx) *UserClaims {
	if u := ctx.Locals("user"); u != nil {
		userClaims, ok := u.(*goJwt.Token).Claims.(*UserClaims)
		if ok {
			return userClaims
		}
	}
	return nil
}

func UserToken(ctx *fiber.Ctx) string {
	if u := ctx.Locals("user"); u != nil {
		return u.(*goJwt.Token).Raw
	}
	return ""
}
