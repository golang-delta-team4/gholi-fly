package http

import (
	"vehicle/internal/vehicle/port"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(vehicleService port.VehicleService) *fiber.App {
	app := fiber.New()

	// Register routes
	RegisterVehicleRoutes(app, vehicleService)

	return app
}
