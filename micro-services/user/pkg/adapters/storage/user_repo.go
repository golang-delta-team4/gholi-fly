package storage

import (
	"context"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPort.Repo {
	return &userRepo{db: db}
}

func (ur userRepo) Create(ctx context.Context, user types.User) error {
	return ur.db.Create(&user).Error
}