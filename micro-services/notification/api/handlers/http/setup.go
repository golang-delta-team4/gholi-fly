package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"

	"notification-nats/config"
)

// RunFiber sets up and starts a Fiber HTTP server.
func RunFiber(db *gorm.DB, cfg config.Config) error {
	router := fiber.New()
	router.Use(recover.New())
	router.Use(logger.New())

	handler := &Handler{DB: db}
	api := router.Group("/api/v1/notification")
	api.Get("/:user_id", handler.GetNotificationsByUserID)

	fiberPort := cfg.Server.HttpPort
	addr := fmt.Sprintf(":%d", fiberPort)
	log.Printf("Starting Fiber on %s...\n", addr)
	return router.Listen(addr)
}
