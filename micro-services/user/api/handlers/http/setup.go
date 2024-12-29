package http

import (
	"fmt"
	"net/http"
	"user-service/api/handlers/shared"
	"user-service/app"
	"user-service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.Config) error {
	userServiceGetter := shared.UserServiceGetter(appContainer, cfg.Server)
	permissionServiceGetter := shared.PermissionServiceGetter(appContainer)
	roleServiceGetter := shared.RoleServiceGetter(appContainer)
	app := fiber.New()
	api := app.Group("api/v1", setUserContext)
	api.Get("health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusAccepted).JSON("Hello World")
	})
	userGroup := api.Group("users")
	userGroup.Get("", setTransaction(appContainer.DB()), GetAllUsers(userServiceGetter))
	userGroup.Post("/sign-up", setTransaction(appContainer.DB()), SignUp(userServiceGetter))
	userGroup.Get("/me", newAuthMiddleware([]byte(cfg.Server.Secret)), GetMe(userServiceGetter))
	userGroup.Post("/sign-in", SignIn(userServiceGetter))
	userGroup.Post("/refresh", newAuthMiddleware([]byte(cfg.Server.Secret)), Refresh(userServiceGetter))
	permissionGroup := api.Group("permissions")
	permissionGroup.Post("/", CreatePermission(permissionServiceGetter))
	roleGroup := api.Group("roles")
	roleGroup.Post("/", CreateRole(roleServiceGetter))
	roleGroup.Get("", GetAllRoles(roleServiceGetter))
	roleGroup.Post("/assign", AssignRole(roleServiceGetter))
	return app.Listen(fmt.Sprintf(":%d", cfg.Server.HttpPort))
}
