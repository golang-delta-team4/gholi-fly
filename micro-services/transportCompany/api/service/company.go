package service

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/domain"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/google/uuid"
)

type CompanyService struct {
	svc                   companyPort.Service
	authSecret            string
	expMin, refreshExpMin uint
}

func NewCompanyService(svc companyPort.Service, authSecret string, expMin, refreshExpMin uint) *CompanyService {
	return &CompanyService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
	}
}

var (
	ErrCompanyCreationValidation = company.ErrCompanyCreationValidation
)

func (s *CompanyService) Create(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	ownerId, err := uuid.Parse(req.OwnerId)
	if err != nil {
		return nil, err
	}

	comanyId, err := s.svc.CreateCompany(ctx, domain.Company{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     ownerId,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateCompanyResponse{
		Id: int64(comanyId),
	}, nil
}
