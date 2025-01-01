package mapper

import (
	"user-service/internal/role/domain"
	"user-service/pkg/adapters/storage/types"
)


func RoleDomain2Storage(role domain.Role) *types.Role {
	return &types.Role{
		Name: role.Name,
	}
}

func RoleStorageList2DomainList(roles []types.Role) []domain.Role {
	var domainRoles []domain.Role
	for _, role := range roles {
		domainRoles = append(domainRoles, domain.Role{
			UUID:      role.UUID,
			Name:     role.Name,
			CreatedAt: role.CreatedAt,
			Permissions: PermissionStorageList2DomainList(role.Permissions),
		})
	}
	return domainRoles
}