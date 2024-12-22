package storage

import (
	"context"
	rolePort "user-service/internal/role/port"
	"user-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) rolePort.Repo {
	return &roleRepo{db: db}
}

func (rr *roleRepo) Create(ctx context.Context,role *types.Role) (uint, error) {
	err := rr.db.Create(&role).Error
	if err != nil {
		return 0, err
	}
	return role.ID, nil 
}

func (rr *roleRepo) AssignRole(ctx context.Context, userRole types.UserRole) (error) {
	return rr.db.Model(&types.UserRole{}).Create(&userRole).Error
	
}

func (rr *roleRepo) GetRole(ctx context.Context, roleUUID uuid.UUID) (*types.Role, error) {
	var role types.Role
	err := rr.db.Model(&types.Role{}).Where("uuid = ?", roleUUID).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}