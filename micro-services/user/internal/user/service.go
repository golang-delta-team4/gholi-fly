package user

import (
	"context"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/storage/mapper"

	"github.com/google/uuid"
)

type service struct {
	repo userPort.Repo
}

func NewService(repo userPort.Repo) userPort.Service {
	return &service{
		repo: repo,
	}
}

func (us service) SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error) {
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

