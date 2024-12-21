package app

import (
	"user-service/config"
	permissionPort "user-service/internal/permission/port"
	userPort "user-service/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	UserService() userPort.Service
	PermissionService() permissionPort.Service
}
