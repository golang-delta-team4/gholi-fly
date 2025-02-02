package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/service"
	"github.com/google/uuid"
)

func CreateCompany(svcGetter ServiceGetter[*service.CompanyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.CreateCompanyRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		ownerUUID := c.Locals("UserUUID").(uuid.UUID)
		response, err := svc.Create(c.UserContext(), &req, ownerUUID)

		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetCompanyById(svcGetter ServiceGetter[*service.CompanyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		companyId := c.Params("id")
		response, err := svc.GetCompanyById(c.UserContext(), companyId)
		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func GetByOwnerId(svcGetter ServiceGetter[*service.CompanyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		companyId := c.Params("ownerId")
		response, err := svc.GetByOwnerId(c.UserContext(), companyId)
		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(response)
	}
}

func UpdateCompany(svcGetter ServiceGetter[*service.CompanyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		companyId := c.Params("id")
		var req pb.UpdateCompanyRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err := svc.UpdateCompany(c.UserContext(), &req, companyId)
		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

func DeleteCompany(svcGetter ServiceGetter[*service.CompanyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		companyId := c.Params("id")
		err := svc.DeleteCompany(c.UserContext(), companyId)
		if err != nil {
			if errors.Is(err, service.ErrCompanyCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
