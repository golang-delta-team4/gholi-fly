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

func (ur *userRepo) Create(ctx context.Context, user types.User) error {
	return ur.db.Create(&user).Error
}

func (ur *userRepo) GetByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	err := ur.db.Model(&types.User{}).Where("email = ? and is_verified = true", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) UpdateRefreshToken(ctx context.Context, userID uint, refreshToken string) error {
	return ur.db.Model(&types.User{}).Where("id = ?", userID).Update("refresh_token", refreshToken).Error
}
