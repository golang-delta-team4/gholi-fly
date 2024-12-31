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
	svc companyPort.Service
}

func NewCompanyService(svc companyPort.Service) *CompanyService {
	return &CompanyService{
		svc: svc,
	}
}

var (
	ErrCompanyCreationValidation = company.ErrCompanyCreationValidation
)

func (s *CompanyService) Create(ctx context.Context, req *pb.CreateCompanyRequest, ownerUUID uuid.UUID) (*pb.CreateCompanyResponse, error) {

	companyId, err := s.svc.CreateCompany(ctx, domain.Company{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     ownerUUID,
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

func (s *CompanyService) GetByOwnerId(ctx context.Context, ownerId string) (*pb.GetCompanyResponse, error) {
	ownerUid, err := uuid.Parse(ownerId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}

	company, err := s.svc.GetByOwnerId(ctx, ownerUid)

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

func (s *CompanyService) UpdateCompany(
	ctx context.Context,
	req *pb.UpdateCompanyRequest,
	companyId string) error {

	companyUId, err := uuid.Parse(companyId)
	if err != nil {
		return fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}

	return s.svc.UpdateCompany(ctx, domain.Company{
		Id:          companyUId,
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
	})
}

func (s *CompanyService) DeleteCompany(ctx context.Context, companyId string) error {
	companyUId, err := uuid.Parse(companyId)
	if err != nil {
		return fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}

	return s.svc.DeleteCompany(ctx, companyUId)
}
