package http

import (
	port_p "gholi-fly-maps/internal/paths/port"
	port_t "gholi-fly-maps/internal/terminals/port"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(terminalService port_t.TerminalService, pathService port_p.PathService) *fiber.App {
	app := fiber.New()

	// Register terminal routes
	RegisterTerminalRoutes(app, terminalService)

	// Register path routes
	RegisterPathRoutes(app, pathService)

	return app
}

