package company

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
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

func (s *service) CreateCompany(ctx context.Context, company domain.Company) (domain.CompanyId, error) {
	if err := company.Validate(); err != nil {
		return 0, fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}
	companyId, err := s.repo.Create(ctx, company)
	if err != nil {
		log.Println("error on creating company: ", err.Error())
	}

	return companyId, nil
}
