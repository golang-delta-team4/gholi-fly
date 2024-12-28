package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/service"
)

func CreateTechnicalTeam(svcGetter ServiceGetter[*service.TechnicalTeamService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.CreateTechnicalTeamRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		response, err := svc.CreateTechnicalTeam(c.UserContext(), &req)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetTechnicalTeamById(svcGetter ServiceGetter[*service.TechnicalTeamService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		teamId := c.Params("id")
		response, err := svc.GetTechnicalTeamById(c.UserContext(), teamId)
		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetTechnicalTeams(svcGetter ServiceGetter[*service.TechnicalTeamService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		pageSize, _ := strconv.Atoi(c.Query("page-size"))
		pageNumber, _ := strconv.Atoi(c.Query("page-number"))

		response, err := svc.GetTechnicalTeams(c.UserContext(), pageSize, pageNumber)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func SetTechTeamToTrip(svcGetter ServiceGetter[*service.TechnicalTeamService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		var req pb.SetTechnicalTeamToTripRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err := svc.SetTechTeamToTrip(c.UserContext(), &req)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func AddTechnicalTeamMember(svcGetter ServiceGetter[*service.TechnicalTeamService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.AddTechnicalTeamMemberRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err := svc.AddTechnicalTeamMember(c.UserContext(), &req)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
