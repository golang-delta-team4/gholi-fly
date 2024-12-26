package http

import (
	"gholi-fly-hotel/pkg/context"
	"gholi-fly-hotel/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext(), context.WithLogger(logger.NewLogger())))
	return c.Next()
}

func setTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tx := db.Begin()

		context.SetDB(c.UserContext(), tx, true)

		if err := c.Next(); err != nil {
			context.Rollback(c.UserContext())
			return err
		}

		if c.Response().StatusCode() >= 300 {
			return context.Rollback(c.UserContext())
		}

		if err := context.CommitOrRollback(c.UserContext(), true); err != nil {
			return err
		}

		return nil
	}
}
