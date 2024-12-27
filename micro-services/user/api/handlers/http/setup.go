package http

import (
	"fmt"
	"net/http"
	"user-service/api/service"
	"user-service/app"
	"user-service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.Config) error {
	userService := service.NewUserService(appContainer.UserService(), cfg.Server.AuthExpMinute, cfg.Server.AuthRefreshMinute, cfg.Server.Secret)
	permissionService := service.NewPermissionService(appContainer.PermissionService())
	roleService := service.NewRoleService(appContainer.RoleService())
	app := fiber.New()
	api := app.Group("api/v1")
	api.Get("health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusAccepted).JSON("Hello World")
	})
	userGroup := api.Group("users")
	userGroup.Post("/sign-up", SignUp(userService))
	userGroup.Get("/me", newAuthMiddleware([]byte(cfg.Server.Secret)), GetMe(userService))
	userGroup.Post("/sign-in", SignIn(userService))
	userGroup.Post("/refresh", newAuthMiddleware([]byte(cfg.Server.Secret)), Refresh(userService))
	permissionGroup := api.Group("permissions")
	permissionGroup.Post("/", CreatePermission(permissionService))
	roleGroup := api.Group("roles")
	roleGroup.Post("/", CreateRole(roleService))
	roleGroup.Post("/assign", AssignRole(roleService))
	return app.Listen(fmt.Sprintf(":%d", cfg.Server.HttpPort))
}
