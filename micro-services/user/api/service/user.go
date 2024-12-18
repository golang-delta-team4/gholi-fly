package service

import (
	"context"
	"fmt"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	"user-service/pkg/jwt"
	"user-service/pkg/time"

	goJwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func (us *UserService) SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	return us.service.SignUp(ctx, user)
}

func (us *UserService) SignIn(ctx context.Context, user *domain.UserSignInRequest) (string, string, error) {
	userID, err := us.service.SignIn(ctx, user)
	if err != nil {
		return "", "", err
	}
	accessToken, err := us.createAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := us.createRefreshToken(userID)
	if err != nil {
		return "", "", err
	}
	err = us.service.UpdateUserRefreshToken(ctx, userID, refreshToken)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil

}

func (us *UserService) createAccessToken(userID uint) (string, error) {
	accessToken, err := jwt.CreateToken([]byte(us.secret), &jwt.UserClaims{UserID: userID, RegisteredClaims: goJwt.RegisteredClaims{ExpiresAt: goJwt.NewNumericDate(time.AddMinutes(us.expMin, true))}})
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (us *UserService) createRefreshToken(userID uint) (string, error) {
	refreshToken, err := jwt.CreateToken([]byte(us.secret), &jwt.UserClaims{UserID: userID, RegisteredClaims: goJwt.RegisteredClaims{ExpiresAt: goJwt.NewNumericDate(time.AddMinutes(us.refreshExpMin, true))}})
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
