package http

import (
	"fmt"

	"gholi-fly-hotel/app"
	"gholi-fly-hotel/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "gholi-hotels-api is running",
		})
	})

	return app.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}
