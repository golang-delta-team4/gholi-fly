package company

import (
	"context"
	"errors"
	"fmt"
	"log"

	rolePB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/role"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	port "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	grpcPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrCompanyOnCreate           = errors.New("error on creating new company")
	ErrCompanyCreationValidation = errors.New("validation failed")
	ErrCompanyNotFound = errors.New("company not found")
)

type service struct {
	repo       port.Repo
	roleClient grpcPort.GRPCRoleClient
}

func NewService(repo port.Repo, roleClient grpcPort.GRPCRoleClient) port.Service {
	return &service{
		repo:       repo,
		roleClient: roleClient,
	}
}

func (s *service) CreateCompany(ctx context.Context, company domain.Company) (uuid.UUID, error) {
	if err := company.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}
	companyId, err := s.repo.Create(ctx, company)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
		return uuid.Nil, err
	}
	_, err = s.roleClient.CreateRole(&rolePB.GrantResourceAccessRequest{OwnerUUID: company.OwnerId.String(), Permissions: createCompanyPermissions(companyId), RoleName: fmt.Sprintf("company-%s", companyId.String())})
	if err != nil {
		log.Println("error on creating role: ", err.Error())
	}
	return companyId, nil
}

func (s *service) GetCompanyById(ctx context.Context, companyId uuid.UUID) (*domain.Company, error) {
	company, err := s.repo.GetCompanyById(ctx, companyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCompanyNotFound
		}
		log.Println("error on getting company: ", err.Error())
		return nil, err
	}
	return company, nil
}

func (s *service) GetByOwnerId(ctx context.Context, ownerId uuid.UUID) (*domain.Company, error) {
	company, err := s.repo.GetByOwnerId(ctx, ownerId)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
		return nil, err
	}
	return company, nil
}

func (s *service) UpdateCompany(ctx context.Context, company domain.Company) error {
	if err := company.UpdateValidate(); err != nil {
		return fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}
	err := s.repo.UpdateCompany(ctx, company)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
		return err
	}
	return nil
}

func (s *service) DeleteCompany(ctx context.Context, companyId uuid.UUID) error {
	err := s.repo.DeleteCompany(ctx, companyId)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
		return err
	}
	return nil
}

func createCompanyPermissions(companyUUID uuid.UUID) []*rolePB.ResourcePermission {
	return []*rolePB.ResourcePermission{
		{
			Route:  fmt.Sprintf("/company/%s/trip", companyUUID.String()),
			Method: "POST",
		},
		{
			Route:  fmt.Sprintf("/company/%s/trip/:id", companyUUID.String()),
			Method: "PATCH",
		},
		{
			Route:  fmt.Sprintf("/company/%s/trip/:id", companyUUID.String()),
			Method: "DELETE",
		},
		{
			Route:  fmt.Sprintf("/company/%s", companyUUID.String()),
			Method: "PATCH",
		},
		{
			Route:  fmt.Sprintf("/company/%s", companyUUID.String()),
			Method: "DELETE",
		},
	}
}
