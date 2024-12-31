package technicalTeam

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/port"
	"github.com/google/uuid"
)

var (
	ErrTechnicalCreationValidation = errors.New("error on create new technical team")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, teamDomain domain.TechnicalTeam) (uuid.UUID, error) {
	if err := teamDomain.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("%w %w", ErrTechnicalCreationValidation, err)
	}
	technicalTeamId, err := s.repo.Create(ctx, teamDomain)
	if err != nil {
		log.Println("error on creating technical team: ", err.Error())
		return uuid.Nil, err
	}

	return technicalTeamId, nil
}

func (s *service) GetById(ctx context.Context, teamId uuid.UUID) (*domain.TechnicalTeam, error) {
	team, err := s.repo.GetById(ctx, teamId)
	if err != nil {
		log.Println("error on getting technical team by id: ", err.Error())
		return &domain.TechnicalTeam{}, err
	}

	return team, nil
}

func (s *service) GetAll(ctx context.Context, pageSize int, page int) ([]domain.TechnicalTeam, error) {
	teams, err := s.repo.GetAll(ctx, pageSize, page)
	if err != nil {
		log.Println("error on getting all technical teams: ", err.Error())
		return nil, err
	}

	return teams, nil
}

func (s *service) SetMember(ctx context.Context, teamId uuid.UUID, technicalTeamMember domain.TechnicalTeamMember) error {
	err := s.repo.SetMember(ctx, teamId, technicalTeamMember)
	if err != nil {
		log.Println("error on setting member to technical team: ", err.Error())
		return err
	}

	return nil
}

func (s *service) SetToTrip(ctx context.Context, teamId uuid.UUID, tripId uuid.UUID) error {
	err := s.repo.SetToTrip(ctx, teamId, tripId)
	if err != nil {
		log.Println("error on setting technical team to trip: ", err.Error())
		return err
	}

	return nil
}
