package presenter

import "github.com/google/uuid"

type CreateRoleRequest struct {
	Name  string `json:"name" validate:"required"`
	PermissionUUIDs []uuid.UUID `json:"permissions" validate:"required,min=1,dive,uuid4"`
}

type AssignRoleRequest struct {
	UserUUID  uuid.UUID `json:"userUUID" validate:"required,uuid4"`
	RoleUUIDs []uuid.UUID `json:"roles" validate:"required,min=1,dive,uuid4"`
}