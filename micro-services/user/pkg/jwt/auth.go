package jwt

import (

	goJwt "github.com/golang-jwt/jwt/v5"
)

const UserClaimKey = "User-Claims"

func CreateToken(secret []byte, claims *UserClaims) (string, error) {
	return goJwt.NewWithClaims(goJwt.SigningMethodHS512, claims).SignedString(secret)
}
