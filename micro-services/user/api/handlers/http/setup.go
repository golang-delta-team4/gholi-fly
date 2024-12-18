package http

import (
	"fmt"
	"net/http"
	"user-service/app"
	"user-service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.Config) error {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(http.StatusAccepted).JSON("Hello World")
	})
	return app.Listen(fmt.Sprintf("%s:%d", cfg.Server.HttpHost, cfg.Server.HttpPort))
}
