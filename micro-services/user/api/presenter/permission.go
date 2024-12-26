package presenter

type CreatePermissionRequest struct {
	Route  string `json:"route" validate:"required"`
	Method string `json:"method" validate:"required"`
}
