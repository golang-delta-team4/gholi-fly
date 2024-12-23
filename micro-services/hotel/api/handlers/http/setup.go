package http

import (
	"fmt"

	"gholi-fly-hotel/app"
	"gholi-fly-hotel/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "gholi-hotels-api is running",
		})
	})

	api := router.Group("/", setUserContext)

	registerHotelAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerHotelAPI(appContainer app.App, router fiber.Router) {
	hotelSvcGetter := hotelServiceGetter(appContainer)
	router.Post("/create", setTransaction(appContainer.DB()), CreateHotel(hotelSvcGetter))
}
