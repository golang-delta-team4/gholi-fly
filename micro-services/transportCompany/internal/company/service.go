package company

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/google/uuid"
)

var (
	ErrCompanyOnCreate           = errors.New("error on creating new company")
	ErrCompanyCreationValidation = errors.New("validation failed")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
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

	return companyId, nil
}

func (s *service) GetCompanyById(ctx context.Context, companyId uuid.UUID) (*domain.Company, error) {
	company, err := s.repo.GetCompanyById(ctx, companyId)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
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
