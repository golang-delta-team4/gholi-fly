package storage

import (
	"context"
	permissionPort "user-service/internal/permission/port"
	"user-service/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type permissionRepo struct {
	db *gorm.DB
}

func NewPermissionRepo(db *gorm.DB) permissionPort.Repo {
	return &permissionRepo{db: db}
}

func (pr *permissionRepo) Create(ctx context.Context, permission []types.Permission) error {
	return pr.db.Create(&permission).Error
}

func (pr *permissionRepo) CheckPermissionExistence(ctx context.Context, route string, method string) (bool, error) {
	var total int
	err := pr.db.Model(&types.Permission{}).Select("count(route)").Where("route = ? and method = ?", route, method).Find(&total).Error
	if err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

func (pr *permissionRepo) GetPermissionsByUUID(ctx context.Context, permissionsUUID []types.Permission) ([]types.Permission, error) {
	var permissions []types.Permission
	err := pr.db.Model(&types.Permission{}).Find(&permissions,permissionsUUID).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
