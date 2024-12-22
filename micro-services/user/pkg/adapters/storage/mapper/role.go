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