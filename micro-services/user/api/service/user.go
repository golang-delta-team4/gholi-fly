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

func (us *UserService) AuthorizeUser(ctx context.Context, req presenter.UserAuthorization) (bool, error) {
	return us.service.AuthorizeUser(ctx, &domain.UserAuthorize{UserUUID: req.UserUUID, Route: strings.ToLower(req.Route), Method: strings.ToLower(req.Method)})
}

func (us *UserService) BlockUser(ctx context.Context, userUUID uuid.UUID) (error) {
	return us.service.BlockUser(ctx, userUUID)
}

func (us *UserService) UnBlockUser(ctx context.Context, userUUID uuid.UUID) (error) {
	return us.service.UnBlockUser(ctx, userUUID)
}

func (us *UserService) GetUserByUUID(ctx context.Context, userUUID string) (*domain.User, error) {
	uuid, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, errors.New("user uuid invalid")
	}
	return us.service.GetUserByUUID(ctx, uuid)
}

func (us *UserService) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*domain.User, error) {
	return us.service.GetUserByEmail(ctx, req.UserEmail)
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

func (us *UserService) GetAllUsers(ctx context.Context, query presenter.PaginationQuery) ([]domain.User, error) {
	if query.Page == 0 {
		query.Page = 1
	} 
	if query.Size == 0 {
		query.Size = 10
	}
	return us.service.GetAllUsers(ctx, query)
}
