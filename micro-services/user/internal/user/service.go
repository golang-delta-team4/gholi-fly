package user

import (
	"context"
	"errors"
	"time"
	"user-service/api/presenter"
	"user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	bankClientPort "user-service/pkg/adapters/clients/grpc/port"
	"user-service/pkg/adapters/storage/mapper"
	"user-service/pkg/adapters/storage/types"

	bankPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/bank"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrEmailNotUnique 		   = errors.New("the email address is already registered.")
	ErrUserNotFound            = errors.New("user not found")
	ErrEmailOrPasswordMismatch = errors.New("email or password mismatch")
)

type service struct {
	repo userPort.Repo
	bankClient bankClientPort.GRPCBankClient
}

func NewService(repo userPort.Repo, bankClient bankClientPort.GRPCBankClient) userPort.Service {
	return &service{
		repo: repo,
		bankClient: bankClient,
	}
}

func (us *service) SignUp(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	storageUser := mapper.Domain2Storage(*user)
	storageUser.UUID = uuid.New()
	var err error
	storageUser.Password, err = domain.HashPassword(user.Password)
	if err != nil {
		return uuid.Nil, err
	}
	err = us.repo.Create(ctx, *storageUser)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok { //gorm does not have any sentinel error for this error type
			if pgErr.Code == "23505" {
				return uuid.Nil, ErrEmailNotUnique
			}
		}
		return uuid.Nil, err
	}
	resp, err := us.bankClient.CreateUserWallet(storageUser.UUID.String())
	if err != nil {
		return uuid.Nil, errors.New("failed to create user wallet")
	}
	if resp.Status == bankPB.ResponseStatus_FAILED {
		return uuid.Nil, err
	}
	return storageUser.UUID, nil
}

func (us *service) SignIn(ctx context.Context, userReq *domain.User) (uuid.UUID, error) {
	user, err := us.repo.GetByEmail(ctx, userReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uuid.Nil, ErrUserNotFound
		}
		return uuid.Nil, err
	}
	passwordMatch := domain.HashVerify(user.Password, userReq.Password)
	if !passwordMatch {
		return uuid.Nil, ErrEmailOrPasswordMismatch
	}
	return user.UUID, nil
}

func (us *service) UpdateUserRefreshToken(ctx context.Context, userID uint, refreshToken string, expirationTime time.Time) error {
	existingRefreshToken, err := us.repo.GetRefreshToken(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = us.repo.AddRefreshToken(ctx, &types.RefreshToken{UserID: userID,
				Token:          refreshToken,
				ExpirationTime: expirationTime})
			return err
		}
		return err
	}
	existingRefreshToken.Token = refreshToken
	existingRefreshToken.ExpirationTime = expirationTime
	return us.repo.UpdateRefreshToken(ctx, existingRefreshToken)
}

func (us *service) GetUserRefreshToken(ctx context.Context, userID uint) (string, error) {
	existingRefreshToken, err := us.repo.GetRefreshToken(ctx, userID)
	if err != nil {
		return "", err
	}
	return existingRefreshToken.Token, nil
}

func (us *service) ValidateRefreshToken(ctx context.Context, userID uint, refreshToken string) (bool, error) {
	existingRefreshToken, err := us.repo.GetRefreshToken(ctx, userID)
	if err != nil {
		return false, err
	}
	if time.Now().After(existingRefreshToken.ExpirationTime) {
		return false, nil
	}
	if existingRefreshToken.Token != refreshToken {
		return false, nil
	}
	return true, nil
}

func (us *service) GetUserIDByUUID(ctx context.Context, userUUID uuid.UUID) (uint, error) {
	user, err := us.repo.GetUserByUUID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrUserNotFound
		}
		return 0, err
	}
	return user.ID, nil
}

func (us *service) AuthorizeUser(ctx context.Context, userAuthorization *domain.UserAuthorize) (bool, error) {
	ok, err := us.repo.AuthorizeUser(ctx, &types.UserAuthorization{UserUUID: userAuthorization.UserUUID, Route: userAuthorization.Route, Method: userAuthorization.Method})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return ok, nil
}

func (us *service) GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*domain.User, error) {
	user, err := us.repo.GetUserByUUID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return mapper.Storage2Domain(*user), nil
}

func (us *service) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := us.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return mapper.Storage2Domain(*user), nil
}

func (us *service) GetAllUsers(ctx context.Context, query presenter.PaginationQuery) ([]domain.User, error) {
	users, err := us.repo.GetAllUsers(ctx, query)
	if err != nil {
		return nil, err
	}
	storageUsers := mapper.StorageList2DomainList(users)
	return storageUsers, nil
}
