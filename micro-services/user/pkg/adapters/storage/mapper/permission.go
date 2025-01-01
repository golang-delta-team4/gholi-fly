package mapper

import (
	"user-service/internal/permission/domain"
	"user-service/pkg/adapters/storage/types"
)

func PermissionDomain2Storage(permission domain.Permission) *types.Permission {
	return &types.Permission{
		Route:  permission.Route,
		Method: permission.Method,
		UUID:   permission.UUID,
	}
}

func PermissionStorageList2DomainList(permissions []types.Permission) []domain.Permission {
	var domainPermissions []domain.Permission
	for _, permission := range permissions {
		domainPermissions = append(domainPermissions, domain.Permission{
			UUID:      permission.UUID,
			Route:     permission.Route,
			Method:    permission.Method,
			CreatedAt: permission.CreatedAt,
		})
	}
	return domainPermissions
}
