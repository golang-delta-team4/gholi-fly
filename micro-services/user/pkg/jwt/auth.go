package jwt

import (
	"errors"

	goJwt "github.com/golang-jwt/jwt/v5"
)

const UserClaimKey = "User-Claims"

func CreateToken(secret []byte, claims *UserClaims) (string, error) {
	return goJwt.NewWithClaims(goJwt.SigningMethodHS512, claims).SignedString(secret)
}

func ParseToken(tokenString string, secret []byte) (*UserClaims, error) {
	token, err := goJwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *goJwt.Token) (interface{}, error) {
		return secret, nil
	})

	if token == nil {
		return nil, errors.New("invalid token (nil)")
	}

	var claim *UserClaims
	if token.Claims != nil {
		cc, ok := token.Claims.(*UserClaims)
		if ok {
			claim = cc
		}
	}

	if err != nil {
		return claim, err
	}

	if !token.Valid {
		return claim, errors.New("token is not valid")
	}

	return claim, nil
}