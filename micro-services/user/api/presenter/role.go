package presenter

type CreateRoleRequest struct {
	Name  string `json:"name" validate:"required"`
	PermissionUUIDs []string `json:"permissions" validate:"required,min=1,dive,uuid4"`
}
