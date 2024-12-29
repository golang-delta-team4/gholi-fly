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

// import (
// 	"net/http"

// 	port_p"gholi-fly-maps/internal/paths/port"
// 	port_t"gholi-fly-maps/internal/terminals/port"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// // SetupRouter initializes the HTTP router and registers all routes.
// func SetupRouter(terminalService port_t.TerminalService, pathService port_p.PathService) http.Handler {
// 	// Create a new router instance
// 	r := chi.NewRouter()

// 	// Apply middleware
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	// Register terminal routes
// 	RegisterTerminalRoutes(r, terminalService)

// 	// Register path routes
// 	RegisterPathRoutes(r, pathService)

// 	return r
// }
