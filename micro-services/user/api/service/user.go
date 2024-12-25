package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"user-service/api/pb"
	"user-service/api/presenter"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	"user-service/pkg/jwt"
	timePkg "user-service/pkg/time"

	goJwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidRefreshToken = errors.New("refresh token is invalid")
)

type ErrUserCreationValidation struct {
	details string
}

func (err *ErrUserCreationValidation) Error() string {
	return fmt.Sprintf("validation failed for: %v", err.details)
}

type UserService struct {
	service               userPort.Service
	expMin, refreshExpMin uint
	secret                string
}

func NewUserService(service userPort.Service, expMin, refreshExpMin uint, secret string) *UserService {
	return &UserService{service: service, expMin: expMin, refreshExpMin: refreshExpMin, secret: secret}
}

func (us *UserService) SignUp(ctx context.Context, user *presenter.UserSignUpRequest) (uuid.UUID, error) {
	return us.service.SignUp(ctx, &domain.User{Email: user.Email, Password: user.Password, FirstName: user.FirstName, LastName: user.LastName})
}

func (us *UserService) SignIn(ctx context.Context, user *presenter.UserSignInRequest) (string, string, error) {
	userUUID, err := us.service.SignIn(ctx, &domain.User{Email: user.Email, Password: user.Password})
	if err != nil {
		return "", "", err
	}
	userID, err := us.service.GetUserIDByUUID(ctx, userUUID)
	if err != nil {
		return "", "", err
	}
	accessToken, _, err := createToken(userUUID, us.expMin, []byte(us.secret))
	if err != nil {
		return "", "", err
	}
	refreshToken, expirationTime, err := createToken(userUUID, us.refreshExpMin, []byte(us.secret))
	if err != nil {
		return "", "", err
	}
	err = us.service.UpdateUserRefreshToken(ctx, userID, refreshToken, expirationTime)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil

}

func (us *UserService) Refresh(ctx context.Context, userUUID uuid.UUID, refreshToken string) (string, string, error) {
	userID, err := us.service.GetUserIDByUUID(ctx, userUUID)
	if err != nil {
		return "", "", err
	}
	valid, err := us.service.ValidateRefreshToken(ctx, userID, refreshToken)
	if err != nil {
		return "", "", err
	}
	if !valid {
		return "", "", ErrInvalidRefreshToken
	}

	accessToken, _, err := createToken(userUUID, us.expMin, []byte(us.secret))
	if err != nil {
		return "", "", err
	}
	refreshToken, expirationTime, err := createToken(userUUID, us.expMin, []byte(us.secret))
	if err != nil {
		return "", "", err
	}
	err = us.service.UpdateUserRefreshToken(ctx, userID, refreshToken, expirationTime)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil

}

func (us *UserService) AuthorizeUser(ctx context.Context, req *pb.UserAuthorizationRequest) (bool, error) {
	uuid, err := uuid.Parse(req.UserUUID)
	if err != nil {
		return false, errors.New("user uuid invalid")
	}
	return us.service.AuthorizeUser(ctx, &domain.UserAuthorize{UserUUID: uuid, Route: strings.ToLower(req.Route), Method: strings.ToLower(req.Method)})
}

func createToken(userUUID uuid.UUID, expMin uint, secret []byte) (string, time.Time, error) {
	expirationTime := timePkg.AddMinutes(expMin, true)
	token, err := jwt.CreateToken(secret, &jwt.UserClaims{UserUUID: userUUID,
		RegisteredClaims: goJwt.RegisteredClaims{ExpiresAt: goJwt.NewNumericDate(expirationTime)}})
	if err != nil {
		return "", time.Time{}, err
	}
	return token, expirationTime, nil
}
