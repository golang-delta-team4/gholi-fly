package storage

import (
	"context"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
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

func (ur *userRepo) UpdateRefreshToken(ctx context.Context, refreshToken types.RefreshToken) error {
	return ur.db.Model(&types.RefreshToken{}).Where("user_id = ?", refreshToken.UserID).Save(&refreshToken).Error
}

func (ur *userRepo) AddRefreshToken(ctx context.Context, userRefreshToken *types.RefreshToken) error {
	return ur.db.Model(&types.RefreshToken{}).Create(&userRefreshToken).Error
}

func (ur *userRepo) DeleteRefreshToken(ctx context.Context, userID uint) error {
	return ur.db.Model(&types.RefreshToken{}).Where("user_id = ?", userID).Delete(&types.RefreshToken{}).Error
}

func (ur *userRepo) GetRefreshToken(ctx context.Context, userID uint) (types.RefreshToken, error) {
	var refreshToken types.RefreshToken
	err := ur.db.Model(&types.RefreshToken{}).Where("user_id = ?", userID).First(&refreshToken).Error
	return refreshToken, err
}

func (ur *userRepo) GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*types.User, error) {
	var user types.User
	err := ur.db.Model(&types.User{}).Where("uuid = ?", userUUID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
