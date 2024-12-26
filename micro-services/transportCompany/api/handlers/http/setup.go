package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/app"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()
	router.Use(recover.New())

	api := router.Group(
		"/api/v1/transport-company",
		setUserContext,
		//newAuthMiddleware([]byte(cfg.Secret)),
	)

	registerCompanyAPI(appContainer, cfg, api)
	registerTripApi(appContainer, cfg, api)
	registerTicketApi(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerCompanyAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	companyServiceGetter := companyServiceGetter(appContainer, cfg)
	router.Post("/company", setTransaction(appContainer.DB()), CreateCompany(companyServiceGetter))
	router.Get("/company/:id", setTransaction(appContainer.DB()), GetCompanyById(companyServiceGetter))
	router.Get("/get-company-by-ownerid/:ownerId", setTransaction(appContainer.DB()), GetByOwnerId(companyServiceGetter))
	router.Patch("/company/:id", setTransaction(appContainer.DB()), UpdateCompany(companyServiceGetter))
	router.Delete("/company/:id", setTransaction(appContainer.DB()), DeleteCompany(companyServiceGetter))
}

func registerTripApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	tripServiceGetter := tripServiceGetter(appContainer, cfg)
	router.Post("/trip", setTransaction(appContainer.DB()), CreateTrip(tripServiceGetter))
	router.Get("/trip/:id", setTransaction(appContainer.DB()), GetTripById(tripServiceGetter))
	router.Get("/agency-trip/:id", setTransaction(appContainer.DB()), GetAgencyTripById(tripServiceGetter))
	router.Get("/trip", setTransaction(appContainer.DB()), GetTrips(tripServiceGetter))
	router.Get("/agency-trip", setTransaction(appContainer.DB()), GetAgencyTrips(tripServiceGetter))
	router.Patch("/trip/:id", setTransaction(appContainer.DB()), UpdateTrip(tripServiceGetter))
	router.Delete("/trip/:id", setTransaction(appContainer.DB()), DeleteTrip(tripServiceGetter))
}

func registerTicketApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	ticketServiceGetter := ticketServiceGetter(appContainer, cfg)
	router.Post("/buy", setTransaction(appContainer.DB()), BuyTicket(ticketServiceGetter))
	// router.Post("/ticket", setTransaction(appContainer.DB()), CreateTicket(ticketServiceGetter))
	// router.Get("/ticket/:id", setTransaction(appContainer.DB()), GetTicketById(ticketServiceGetter))
	// router.Get("/ticket", setTransaction(appContainer.DB()), GetTickets(ticketServiceGetter))
	// router.Patch("/ticket/:id", setTransaction(appContainer.DB()), UpdateTicket(ticketServiceGetter))
	// router.Delete("/ticket/:id", setTransaction(appContainer.DB()), DeleteTicket(ticketServiceGetter))
}
