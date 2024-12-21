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
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(http.StatusAccepted).JSON("Hello World")
	})
	userGroup := app.Group("users")
	userGroup.Post("/sign-up", SignUp(userService))
	userGroup.Post("/sign-in", SignIn(userService))
	userGroup.Post("/refresh", newAuthMiddleware([]byte(cfg.Server.Secret)), Refresh(userService))
	permissionGroup := app.Group("permissions")
	permissionGroup.Post("/", CreatePermission(permissionService))
	return app.Listen(fmt.Sprintf("%s:%d", cfg.Server.HttpHost, cfg.Server.HttpPort))
}
