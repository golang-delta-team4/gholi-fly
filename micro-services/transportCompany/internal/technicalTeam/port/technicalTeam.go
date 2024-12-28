package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/domain"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, technicalTeam domain.TechnicalTeam) (uuid.UUID, error)
	GetById(ctx context.Context, technicalTeamId uuid.UUID) (*domain.TechnicalTeam, error)
	GetAll(ctx context.Context, pageSize int, page int) ([]domain.TechnicalTeam, error)
	SetMember(ctx context.Context, teamId uuid.UUID, technicalTeamMember domain.TechnicalTeamMember) error
	SetToTrip(ctx context.Context, teamId uuid.UUID, tripId uuid.UUID) error
}
