package shared

import (
	"context"

	"user-service/api/service"
	"user-service/app"
	"user-service/config"
)

func RoleServiceGetter(appContainer app.App) ServiceGetter[*service.RoleService] {
	return func(ctx context.Context) *service.RoleService {
		return service.NewRoleService(appContainer.RoleService(ctx))
	}
}

func PermissionServiceGetter(appContainer app.App) ServiceGetter[*service.PermissionService] {
	return func(ctx context.Context) *service.PermissionService {
		return service.NewPermissionService(appContainer.PermissionService(ctx))
	}
}

func UserServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.UserService] {
	return func(ctx context.Context) *service.UserService {
		return service.NewUserService(appContainer.UserService(ctx), cfg.AuthExpMinute, cfg.AuthRefreshMinute, cfg.Secret)
	}
}
