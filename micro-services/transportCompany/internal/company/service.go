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
