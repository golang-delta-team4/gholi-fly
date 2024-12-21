package service

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}

	companyId, err := s.svc.CreateCompany(ctx, domain.Company{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     ownerId,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateCompanyResponse{
		Id: companyId.String(),
	}, nil
}

func (s *CompanyService) GetCompanyById(ctx context.Context, companyId string) (*pb.GetCompanyResponse, error) {
	companyUid, err := uuid.Parse(companyId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}

	company, err := s.svc.GetCompanyById(ctx, companyUid)

	if err != nil {
		return nil, err
	}

	return &pb.GetCompanyResponse{
		Id:          company.Id.String(),
		Name:        company.Name,
		Description: company.Description,
		Address:     company.Address,
		Phone:       company.Phone,
		Email:       company.Email,
		OwnerId:     company.OwnerId.String(),
	}, nil
}
