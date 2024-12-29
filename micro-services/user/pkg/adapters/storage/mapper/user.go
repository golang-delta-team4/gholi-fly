package mapper

import (
	"user-service/internal/user/domain"
	roleDomain "user-service/internal/role/domain"
	"user-service/pkg/adapters/storage/types"
)

func Domain2Storage(user domain.User) *types.User {
	return &types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func Storage2Domain(user types.User) *domain.User {
	return &domain.User{
		UUID:      user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func StorageList2DomainList(users []types.User) []domain.User {
	var domainUsers []domain.User
	for _, user := range users {
		var roles []roleDomain.Role
		for _, role := range user.UserRoles {
			roles = append(roles, roleDomain.Role{Name: role.Role.Name, UUID: role.Role.UUID})
		}
		domainUsers = append(domainUsers, domain.User{
			UUID:      user.UUID,
			FirstName: user.LastName,
			LastName:  user.FirstName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Roles: roles,
		})
	}
	return domainUsers
}
