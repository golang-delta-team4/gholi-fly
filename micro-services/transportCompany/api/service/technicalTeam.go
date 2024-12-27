package service

import (
	"context"
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/domain"
	technicalTeamPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/port"
	"github.com/google/uuid"
)

var (
	ErrTechnicalCreationValidation = technicalTeam.ErrTechnicalCreationValidation
)

type TechnicalTeamService struct {
	svc technicalTeamPort.Service
}

func NewTechnicalTeamService(svc technicalTeamPort.Service) *TechnicalTeamService {
	return &TechnicalTeamService{
		svc: svc,
	}
}

func (s *TechnicalTeamService) CreateTechnicalTeam(ctx context.Context, req *pb.CreateTechnicalTeamRequest) (*pb.CreateTechnicalTeamResponse, error) {
	companyId, err := uuid.Parse(req.CompanyId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTechnicalCreationValidation, err)
	}

	technicalTeamId, err := s.svc.Create(ctx, domain.TechnicalTeam{
		Name:        req.Name,
		CompanyId:   companyId,
		Description: req.Description,
		TripType:    req.TripType,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateTechnicalTeamResponse{
		Id: technicalTeamId.String(),
	}, nil
}

func (s *TechnicalTeamService) GetTechnicalTeamById(ctx context.Context, technicalTeamId string) (*pb.GetTechnicalTeamResponse, error) {
	teamId, err := uuid.Parse(technicalTeamId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrTechnicalCreationValidation, err)
	}

	team, err := s.svc.GetById(ctx, teamId)

	if err != nil {
		return nil, err
	}

	return &pb.GetTechnicalTeamResponse{
		Id:          team.Id.String(),
		Name:        team.Name,
		Description: team.Description,
		TripType:    team.TripType,
		CompanyId:   team.CompanyId.String(),
		MembersId:   team.Members,
	}, nil
}
