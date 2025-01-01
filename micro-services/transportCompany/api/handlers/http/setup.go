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
	)

	registerCompanyAPI(appContainer, cfg, api)
	registerTripApi(appContainer, cfg, api)
	registerTicketApi(appContainer, cfg, api)
	registerTechnicalTeamApi(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerCompanyAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	companyServiceGetter := companyServiceGetter(appContainer, cfg)

	router.Post("/company", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), CreateCompany(companyServiceGetter))
	router.Get("/company/:id", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetCompanyById(companyServiceGetter))
	router.Get("/get-company-by-ownerid/:ownerId", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetByOwnerId(companyServiceGetter))
	router.Patch("/company/:id", newAuthMiddleware([]byte(cfg.Secret)), newAuthorizationMiddlewareDirect(appContainer.UserGRPCService()), setTransaction(appContainer.DB()), UpdateCompany(companyServiceGetter))
	router.Delete("/company/:id", newAuthMiddleware([]byte(cfg.Secret)), newAuthorizationMiddlewareDirect(appContainer.UserGRPCService()), setTransaction(appContainer.DB()), DeleteCompany(companyServiceGetter))
}

func registerTripApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	tripServiceGetter := tripServiceGetter(appContainer, cfg)
	router.Post("/trip", setTransaction(appContainer.DB()), CreateTrip(tripServiceGetter, appContainer.UserGRPCService()))
	router.Get("/trip/:id", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetTripById(tripServiceGetter))
	router.Get("/agency-trip/:id", setTransaction(appContainer.DB()), GetAgencyTripById(tripServiceGetter))
	router.Get("/trip", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetTrips(tripServiceGetter))
	router.Get("/agency-trip", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetAgencyTrips(tripServiceGetter))
	router.Patch("/trip/:id", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), UpdateTrip(tripServiceGetter, appContainer.UserGRPCService()))
	router.Delete("/trip/:id", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), DeleteTrip(tripServiceGetter, appContainer.UserGRPCService()))

	router.Patch("/confirm-trip/:id", setTransaction(appContainer.DB()), ConfirmTrip(tripServiceGetter))
}

func registerTicketApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	ticketServiceGetter := ticketServiceGetter(appContainer, cfg)
	router.Post("/ticket/buy", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), BuyTicket(ticketServiceGetter))
	router.Post("/ticket/agency-buy", setTransaction(appContainer.DB()), BuyAgencyTicket(ticketServiceGetter))
	router.Post("/ticket/cancel/:id", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), CancelTicket(ticketServiceGetter))
	router.Post("/agency-ticket/cancel/:id", setTransaction(appContainer.DB()), CancelAgencyTicket(ticketServiceGetter))
}

func registerTechnicalTeamApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	technicalTeamServiceGetter := technicalTeamServiceGetter(appContainer, cfg)
	router.Post("/technical-team", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), CreateTechnicalTeam(technicalTeamServiceGetter))
	router.Post("/technical-team-member", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), AddTechnicalTeamMember(technicalTeamServiceGetter))
	router.Get("/technical-team/:id", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetTechnicalTeamById(technicalTeamServiceGetter))
	router.Get("/technical-team", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), GetTechnicalTeams(technicalTeamServiceGetter))
	router.Patch("/set-technical-team/", newAuthMiddleware([]byte(cfg.Secret)), setTransaction(appContainer.DB()), SetTechTeamToTrip(technicalTeamServiceGetter))
}
