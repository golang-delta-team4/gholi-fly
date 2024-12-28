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
		newAuthMiddleware([]byte(cfg.Secret)),
	)

	registerCompanyAPI(appContainer, cfg, api)
	registerTripApi(appContainer, cfg, api)
	registerTicketApi(appContainer, cfg, api)
	registerTechnicalTeamApi(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerCompanyAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	companyServiceGetter := companyServiceGetter(appContainer, cfg)
	router.Post("/company", setTransaction(appContainer.DB()), CreateCompany(companyServiceGetter))
	router.Get("/company/:id", setTransaction(appContainer.DB()), GetCompanyById(companyServiceGetter))
	router.Get("/get-company-by-ownerid/:ownerId", setTransaction(appContainer.DB()), GetByOwnerId(companyServiceGetter))
	router.Patch("/company/:id", newAuthorizationMiddlewareDirect(appContainer.UserGRPCService()), setTransaction(appContainer.DB()), UpdateCompany(companyServiceGetter))
	router.Delete("/company/:id", newAuthorizationMiddlewareDirect(appContainer.UserGRPCService()), setTransaction(appContainer.DB()), DeleteCompany(companyServiceGetter))
}

func registerTripApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	tripServiceGetter := tripServiceGetter(appContainer, cfg)
	router.Post("/trip", setTransaction(appContainer.DB()), CreateTrip(tripServiceGetter, appContainer.UserGRPCService()))
	router.Get("/trip/:id", setTransaction(appContainer.DB()), GetTripById(tripServiceGetter))
	router.Get("/agency-trip/:id", setTransaction(appContainer.DB()), GetAgencyTripById(tripServiceGetter))
	router.Get("/trip", setTransaction(appContainer.DB()), GetTrips(tripServiceGetter))
	router.Get("/agency-trip", setTransaction(appContainer.DB()), GetAgencyTrips(tripServiceGetter))
	router.Patch("/trip/:id", setTransaction(appContainer.DB()), UpdateTrip(tripServiceGetter, appContainer.UserGRPCService()))
	router.Delete("/trip/:id", setTransaction(appContainer.DB()), DeleteTrip(tripServiceGetter))

	// router.Patch("/cancel-trip/:id", setTransaction(appContainer.DB()), CancelTrip(tripServiceGetter))
	// router.Patch("/finish-trip/:id", setTransaction(appContainer.DB()), FinishTrip(tripServiceGetter))
	// router.Patch("/confirm-trip/:id", setTransaction(appContainer.DB()), ConfirmTrip(tripServiceGetter))
}

func registerTicketApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	ticketServiceGetter := ticketServiceGetter(appContainer, cfg)
	router.Post("/ticket/buy", setTransaction(appContainer.DB()), BuyTicket(ticketServiceGetter))
	router.Post("/ticket/agency-buy", setTransaction(appContainer.DB()), BuyAgencyTicket(ticketServiceGetter))
	router.Post("/ticket/cancel/:id", setTransaction(appContainer.DB()), CancelTicket(ticketServiceGetter))
}

func registerTechnicalTeamApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	technicalTeamServiceGetter := technicalTeamServiceGetter(appContainer, cfg)
	router.Post("/technical-team", setTransaction(appContainer.DB()), CreateTechnicalTeam(technicalTeamServiceGetter))
	router.Post("/technical-team-member", setTransaction(appContainer.DB()), AddTechnicalTeamMember(technicalTeamServiceGetter))
	router.Get("/technical-team/:id", setTransaction(appContainer.DB()), GetTechnicalTeamById(technicalTeamServiceGetter))
	router.Get("/technical-team", setTransaction(appContainer.DB()), GetTechnicalTeams(technicalTeamServiceGetter))
	router.Patch("/set-technical-team/", setTransaction(appContainer.DB()), SetTechTeamToTrip(technicalTeamServiceGetter))
}
