package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"gholi-fly-agancy/app"
	"gholi-fly-agancy/config"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	// Initialize Fiber router
	router := fiber.New()
	router.Use(recover.New())

	// Apply global middlewares
	api := router.Group(
		"/api/v1/agency",
		setUserContext,
		newAuthMiddleware([]byte(cfg.Secret)),
	)

	// Register Agency API
	registerAgencyAPI(appContainer, cfg, api)

	// Start the server
	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAgencyAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	agencyServiceGetter := agencyServiceGetter(appContainer)

	// Create agency (transactional)
	router.Post("/", setTransaction(appContainer.DB()), CreateAgency(agencyServiceGetter))

	// Get agency by ID (non-transactional)
	router.Get("/:id", GetAgency(agencyServiceGetter))

	// Update agency by ID (transactional)
	router.Patch("/:id", setTransaction(appContainer.DB()), UpdateAgency(agencyServiceGetter))

	// Delete agency by ID (non-transactional)
	router.Delete("/:id", DeleteAgency(agencyServiceGetter))
}
