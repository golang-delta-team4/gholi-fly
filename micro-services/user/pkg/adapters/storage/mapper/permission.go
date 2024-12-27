package mapper

import (
	"user-service/internal/permission/domain"
	"user-service/pkg/adapters/storage/types"
)

func PermissionDomain2Storage(permission domain.Permission) *types.Permission {
	return &types.Permission{
		Route:  permission.Route,
		Method: permission.Method,
		UUID: permission.UUID,
	}
}
