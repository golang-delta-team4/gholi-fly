package app

import (
	"context"
	"user-service/config"
	permissionPort "user-service/internal/permission/port"
	rolePort "user-service/internal/role/port"
	userPort "user-service/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	UserService(ctx context.Context) userPort.Service
	PermissionService(ctx context.Context) permissionPort.Service
	RoleService(ctx context.Context) rolePort.Service
}
