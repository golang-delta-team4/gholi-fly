package http

import (
	"vehicle/app"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(appContainer app.App) *fiber.App {
	app := fiber.New()

	// Register routes
	RegisterVehicleRoutes(app, appContainer.VehicleService(), appContainer.Config())

	return app
}
