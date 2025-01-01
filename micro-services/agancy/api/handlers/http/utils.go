package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ServiceGetter[T any] func(context.Context) T

func GetLogger(c *fiber.Ctx) *zap.Logger {
	logger, ok := c.UserContext().Value("logger").(*zap.Logger)
	if !ok {
		panic("Logger not found in context") // Handle gracefully in production
	}
	return logger
}
