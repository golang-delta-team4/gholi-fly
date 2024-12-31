package http

import (
	"fmt"

	"gholi-fly-bank/app"
	"gholi-fly-bank/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()
	router.Use(recover.New())
	api := router.Group("/api/v1/bank", setUserContext)
	api.Use(newAuthMiddleware([]byte(cfg.Secret)))
	registerGlobalRoutes(api)

	registerWalletAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerGlobalRoutes(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "gholi-fly-bank-api is running",
		})
	})
}

func registerWalletAPI(appContainer app.App, router fiber.Router) {
	walletSvcGetter := walletServiceGetter(appContainer)
	router.Get("/wallets", GetWallets(walletSvcGetter))
	router.Patch("/wallets/:id", setTransaction(appContainer.DB()), UpdateWalletBalance(walletSvcGetter))
}
