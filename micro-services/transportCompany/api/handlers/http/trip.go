package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/service"
	clientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"
	"github.com/google/uuid"
)

func CreateTrip(svcGetter ServiceGetter[*service.TripService], userGRPCService clientPort.GRPCUserClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.CreateTripRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err := authorization(userGRPCService, c.Path(), c.Method(), req.CompanyId, c.Locals("UserUUID").(uuid.UUID))
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}
		response, err := svc.CreateTrip(c.UserContext(), &req)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetTripById(svcGetter ServiceGetter[*service.TripService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		tripId := c.Params("id")
		response, err := svc.GetTripById(c.UserContext(), tripId)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetAgencyTripById(svcGetter ServiceGetter[*service.TripService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		tripId := c.Params("id")
		response, err := svc.GetAgencyTripById(c.UserContext(), tripId)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetTrips(svcGetter ServiceGetter[*service.TripService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		pageSize, _ := strconv.Atoi(c.Query("page-size"))
		pageNumber, _ := strconv.Atoi(c.Query("page-number"))

		response, err := svc.GetTrips(c.UserContext(), pageSize, pageNumber)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetAgencyTrips(svcGetter ServiceGetter[*service.TripService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		pageSize, _ := strconv.Atoi(c.Query("page-size"))
		pageNumber, _ := strconv.Atoi(c.Query("page-number"))
		response, err := svc.GetAgencyTrips(c.UserContext(), pageSize, pageNumber)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func UpdateTrip(svcGetter ServiceGetter[*service.TripService], userGRPCService clientPort.GRPCUserClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		tripId := c.Params("id")
		var req pb.UpdateTripRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err := authorization(userGRPCService, c.Path(), c.Method(), req.CompanyId, c.Locals("UserUUID").(uuid.UUID))
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}
		err = svc.UpdateTrip(c.UserContext(), tripId, &req)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

func DeleteTrip(svcGetter ServiceGetter[*service.TripService]) fiber.Handler { //add authorization
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		tripId := c.Params("id")
		err := svc.DeleteTrip(c.UserContext(), tripId)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
