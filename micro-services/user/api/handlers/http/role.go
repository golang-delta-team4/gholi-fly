package http

import (
	"errors"
	"user-service/api/handlers/shared"
	"user-service/api/presenter"
	"user-service/api/service"
	"user-service/internal/role"
	"user-service/internal/user"
	"user-service/pkg/adapters/storage/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateRole(svcGetter shared.ServiceGetter[*service.RoleService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req presenter.CreateRoleRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		resp, err := svc.Create(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, &service.ErrInvalidUUID{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			if errors.Is(err, role.ErrRoleNameNotUnique) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(resp)
	}
}

func AssignRole(svcGetter shared.ServiceGetter[*service.RoleService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req presenter.AssignRoleRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		err := svc.Assign(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, &role.ErrRoleNotFound{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			if errors.Is(err, user.ErrUserNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON("saved changes")
	}
}

func GetAllRoles(svcGetter shared.ServiceGetter[*service.RoleService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req presenter.PaginationQuery
		if err := c.QueryParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		validationError := validate(req)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationError})
		}
		resp, err := svc.GetAllRoles(c.UserContext(), req)
		if err != nil {
			if errors.Is(err, &service.ErrInvalidUUID{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			if errors.Is(err, role.ErrRoleNameNotUnique) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(fiber.Map{"rolesList": resp})
	}
}

func DeleteRole(svcGetter shared.ServiceGetter[*service.RoleService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id := c.Params("id")
		roleUUID, err := uuid.Parse(id)
		if err != nil {
			return fiber.ErrBadRequest
		}
		err = svc.DeleteRole(c.UserContext(), roleUUID)
		if err != nil {
			if errors.Is(err, &role.ErrRoleNotFound{}) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			if errors.Is(err, types.ErrUnableToDeleteSuperAdmin) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())	
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON("role deleted")
	}
}
