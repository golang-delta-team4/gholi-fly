package storage

import (
	"context"
	"user-service/api/presenter"
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
	err := ur.db.Model(&types.User{}).Where("email = ?", email).First(&user).Error
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

func (ur *userRepo) AuthorizeUser(ctx context.Context, userAuthorization *types.UserAuthorization) (bool, error) {
	var total int
	err := ur.db.Model(&types.User{}).Debug().
		Joins("left join user_roles ur on users.id = ur.user_id").
		Joins("left join roles r on r.id = ur.role_id").
		Joins("left join role_permissions rp on rp.role_id = r.id").
		Joins("left join permissions p on rp.permission_id = p.id").
		Where("users.is_blocked = false and users.uuid = ? and ((? like replace(p.route, ':id', '%') and p.method = ?) or r.name = 'SuperAdmin')", userAuthorization.UserUUID, userAuthorization.Route, userAuthorization.Method).
		Select("count(users.id)").Find(&total).Error

	if err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil

}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	err := ur.db.Model(&types.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) GetAllUsers(ctx context.Context, query presenter.PaginationQuery) ([]types.User, error) {
		var users []types.User
		err := ur.db.Model(&types.User{}).Preload("UserRoles.Role").Limit(query.Size).Offset((query.Page-1)*query.Page).Where("deleted_at is null").Find(&users).Error
		if err != nil {
			return nil, err
		}
		return users, nil
}

func (ur *userRepo) Block(ctx context.Context, userUUID uuid.UUID) error {
	return ur.db.Model(&types.User{}).Where("uuid = ?", userUUID).Update("is_blocked", true).Error
}
func (ur *userRepo) UnBlock(ctx context.Context, userUUID uuid.UUID) error {
	return ur.db.Model(&types.User{}).Where("uuid = ?", userUUID).Update("is_blocked", false).Error

}

func (ur *userRepo) GetBlocked(ctx context.Context) ([]uuid.UUID, error) {
	var uuids []uuid.UUID
	err := ur.db.Model(&types.User{}).Select("uuid").Where("is_blocked = true").Find(&uuids).Error
	return uuids, err
}
