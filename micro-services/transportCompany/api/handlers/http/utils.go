package http

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/user"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/jwt"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	clientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

func userClaims(ctx *fiber.Ctx) *jwt.UserClaims {
	if u := ctx.Locals("user"); u != nil {
		userClaims, ok := u.(*jwt2.Token).Claims.(*jwt.UserClaims)
		if ok {
			return userClaims
		}
	}
	return nil
}

type ServiceGetter[T any] func(context.Context) T

func authorization(userGRPCService clientPort.GRPCUserClient, route string, method string, companyID string, userUUID uuid.UUID) error {

	routeDetail := strings.Split(route, "/")
	accessRoute := fmt.Sprintf("/company/%v/%v",companyID,strings.Join(routeDetail[4:], "/"))
	resp, err := userGRPCService.CheckUserAuthorization(&pb.UserAuthorizationRequest{UserUUID: userUUID.String(), Route: accessRoute,Method: method})
	if err != nil {
		return err
	}
	if resp.AuthorizationStatus == pb.Status_FAILED {
		return fiber.ErrForbidden
	}
	return nil
}
