package user

import (
	"context"
	"errors"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/storage/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrEmailOrPasswordMismatch struct{}

func (err ErrEmailOrPasswordMismatch) Error() string {
	return "email or password mismatch"
}

type ErrUserNotFound struct{}

func (err ErrUserNotFound) Error() string {
	return "user not found"
}

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
			return 0, ErrUserNotFound{}
		}
		return 0, err
	}
	passwordMatch := domain.HashVerify(user.Password, userSingInRequest.Password)
	if !passwordMatch {
		return 0, ErrEmailOrPasswordMismatch{}
	}
	return user.ID, nil
}

func (us *service) UpdateUserRefreshToken (ctx context.Context, userID uint, refreshToken string) error {
	return us.repo.UpdateRefreshToken(ctx, userID, refreshToken)
}

