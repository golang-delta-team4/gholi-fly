package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"gholi-fly-agancy/app"
	"gholi-fly-agancy/config"
)

func Run(appContainer app.App, cfg config.Config) error {
	// Initialize Fiber router
	router := fiber.New()
	router.Use(recover.New())
	// Initialize and use LoggerMiddleware
	loggerMiddleware, err := LoggerMiddleware(cfg.Logger)
	if err != nil {
		return err
	}
	router.Use(loggerMiddleware)
	// Apply global middlewares
	api := router.Group(
		"/api/v1/agency",
		setUserContext,
		newAuthMiddleware([]byte(cfg.Server.Secret)),
	)

	// Register Agency API
	registerAgencyAPI(appContainer, api)

	// Register Tour API
	registerTourAPI(appContainer, cfg, api)

	// Register Reservation API
	registerReservationAPI(appContainer, cfg, api)

	// Start the server
	return router.Listen(fmt.Sprintf(":%d", cfg.Server.HttpPort))
}

func registerAgencyAPI(appContainer app.App, router fiber.Router) {
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
func registerTourAPI(appContainer app.App, cfg config.Config, router fiber.Router) {
	tourServiceGetter := tourServiceGetter(appContainer)

	// Create tour (transactional)
	router.Post("/tour/:agencyID", setTransaction(appContainer.DB()), CreateTour(tourServiceGetter, cfg))

	// Get tour by ID (non-transactional)
	router.Get("/tour/:id", GetTour(tourServiceGetter))

	// Update tour by ID (transactional)
	router.Patch("/tour/:id", setTransaction(appContainer.DB()), UpdateTour(tourServiceGetter))

	// Delete tour by ID (non-transactional)
	router.Delete("/tour/:id", DeleteTour(tourServiceGetter))

	// List tours by agency (non-transactional)
	router.Get("/agency/:agencyID/tours", ListToursByAgency(tourServiceGetter))
}
func registerReservationAPI(appContainer app.App, cfg config.Config, router fiber.Router) {
	reservationServiceGetter := reservationServiceGetter(appContainer)

	// Create reservation (transactional) with agency ID
	router.Post("/reservation/:agencyID", setTransaction(appContainer.DB()), CreateReservation(reservationServiceGetter, cfg))
}
