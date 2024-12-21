package http

import (
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/app"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1/transport-company", setUserContext)

	registerCompanyAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerCompanyAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	companyServiceGetter := companyServiceGetter(appContainer, cfg)
	router.Post("/company", setTransaction(appContainer.DB()), CreateCompany(companyServiceGetter))
	router.Get("/company/:id", setTransaction(appContainer.DB()), GetCompanyById(companyServiceGetter))
	router.Get("/get_company-by-ownerid/:ownerId", setTransaction(appContainer.DB()), GetByOwnerId(companyServiceGetter))
	router.Patch("/company/:id", setTransaction(appContainer.DB()), UpdateCompany(companyServiceGetter))
	router.Delete("/company/:id", setTransaction(appContainer.DB()), DeleteCompany(companyServiceGetter))
}
