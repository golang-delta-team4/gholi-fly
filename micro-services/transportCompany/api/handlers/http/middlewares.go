package http

import (
	"strings"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/logger"
	"github.com/google/uuid"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/context"

	"github.com/golang-delta-team4/gholi-fly-shared/jwt"
	pb "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/user"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	clientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"
	"gorm.io/gorm"
)

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			userClaims := jwt.GetUserClaims(ctx)
			if userClaims == nil {
				return fiber.ErrUnauthorized
			}
			ctx.Locals("UserUUID", userClaims.UserUUID)
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
		AuthScheme: "Bearer",
	})
}

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

func newAuthorizationMiddlewareDirect(userGRPCService clientPort.GRPCUserClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userUUID := c.Locals("UserUUID")
		if userUUID == nil {
			return fiber.ErrUnauthorized
		}
		routeDetail := strings.Split(c.Path(), "/")
		resp, err := userGRPCService.CheckUserAuthorization(&pb.UserAuthorizationRequest{UserUUID: userUUID.(uuid.UUID).String(), Route: "/" + strings.Join(routeDetail[4:], "/"), Method: c.Method()})
		if err != nil {
			return err
		}
		if resp.AuthorizationStatus == pb.Status_FAILED {
			return fiber.ErrForbidden
		}
		return c.Next()
	}
}
