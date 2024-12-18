package mapper

import (
	"user-service/internal/user/domain"
	"user-service/pkg/adapters/storage/types"
)

func Domain2Storage(user domain.User) *types.User {
	return &types.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
	}
}

func Storage2Domain(user types.User) *domain.User {
	return &domain.User{
		UUID: user.UUID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}