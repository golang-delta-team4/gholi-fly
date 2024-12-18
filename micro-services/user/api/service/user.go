package service

import (
	"context"
	"fmt"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"

	"github.com/google/uuid"
)

type ErrUserCreationValidation struct {
	details string
}

func (err *ErrUserCreationValidation) Error() string {
	return fmt.Sprintf("validation failed for: %v", err.details)
}

type UserService struct {
	service userPort.Service
	expMin, refreshExpMin uint
}

func NewUserService(service userPort.Service, expMin, refreshExpMin uint) *UserService {
	return &UserService{service: service, expMin: expMin, refreshExpMin: refreshExpMin}
}

func (us *UserService) SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	return us.service.SignUp(ctx, user)
}