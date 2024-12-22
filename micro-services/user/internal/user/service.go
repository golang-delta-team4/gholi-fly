package user

import (
	"context"
	"errors"
	"time"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/storage/mapper"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound            = errors.New("user not found")
	ErrEmailOrPasswordMismatch = errors.New("email or password mismatch")
)

type service struct {
	repo userPort.Repo
}

func NewService(repo userPort.Repo) userPort.Service {
	return &service{
		repo: repo,
	}
}

func (us *service) SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	storageUser := mapper.Domain2Storage(*user)
	storageUser.UUID = uuid.New()
	var err error
	storageUser.Password, err = domain.HashPassword(user.Password)
	if err != nil {
		return uuid.Nil, err
	}
	err = us.repo.Create(ctx, *storageUser)
	if err != nil {
		return uuid.Nil, err
	}
	return storageUser.UUID, nil
}

func (us *service) SignIn(ctx context.Context, userSingInRequest *domain.UserSignInRequest) (uint, error) {

	user, err := us.repo.GetByEmail(ctx, userSingInRequest.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrUserNotFound
		}
		return 0, err
	}
	passwordMatch := domain.HashVerify(user.Password, userSingInRequest.Password)
	if !passwordMatch {
		return 0, ErrEmailOrPasswordMismatch
	}
	return user.ID, nil
}

func (us *service) UpdateUserRefreshToken(ctx context.Context, userID uint, refreshToken string, expirationTime time.Time) error {
	existingRefreshToken, err := us.repo.GetRefreshToken(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = us.repo.AddRefreshToken(ctx, &types.RefreshToken{UserID: userID,
				Token:          refreshToken,
				ExpirationTime: expirationTime})
			return err
		}
		return err
	}
	existingRefreshToken.Token = refreshToken
	existingRefreshToken.ExpirationTime = expirationTime
	return us.repo.UpdateRefreshToken(ctx, existingRefreshToken)
}

func (us *service) GetUserRefreshToken(ctx context.Context, userID uint) (string, error) {
	existingRefreshToken, err := us.repo.GetRefreshToken(ctx, userID)
	if err != nil {
		return "", err
	}
	return existingRefreshToken.Token, nil
}

func (us *service) ValidateRefreshToken(ctx context.Context, userID uint, refreshToken string) (bool, error) {
	existingRefreshToken, err := us.repo.GetRefreshToken(ctx, userID)
	if err != nil {
		return false, err
	}
	if time.Now().After(existingRefreshToken.ExpirationTime) {
		return false, nil
	}
	if existingRefreshToken.Token != refreshToken {
		return false, nil
	}
	return true, nil
}

func (us *service) GetUserIDByUUID(ctx context.Context, userUUID uuid.UUID) (uint, error) {
	user, err := us.repo.GetUserByUUID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrUserNotFound
		}
		return 0, err
	}
	return user.ID, nil
}
