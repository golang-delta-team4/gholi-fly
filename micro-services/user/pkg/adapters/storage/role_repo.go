package storage

import (
	"context"
	rolePort "user-service/internal/role/port"
	"user-service/pkg/adapters/storage/types"

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