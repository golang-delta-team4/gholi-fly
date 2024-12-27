package storage

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/mapper"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type technicalTeamRepo struct {
	db *gorm.DB
}

func NewTechnicalTeamRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	repo := &technicalTeamRepo{db}
	return repo
}

func (r *technicalTeamRepo) Create(ctx context.Context, technicalTeamDomain domain.TechnicalTeam) (uuid.UUID, error) {
	technicalTeam := mapper.TechnicalTeamDomain2Storage(technicalTeamDomain)
	return technicalTeam.Id, r.db.Table("technical_teams").WithContext(ctx).Create(technicalTeam).Error
}
func (r *technicalTeamRepo) GetById(ctx context.Context, technicalTeamId uuid.UUID) (*domain.TechnicalTeam, error) {
	var technicalTeam domain.TechnicalTeam
	err := r.db.Table("technical_teams").WithContext(ctx).Where("id = ?", technicalTeamId).First(&technicalTeam).Error
	if err != nil {
		return nil, err
	}
	return &technicalTeam, nil
}
func (r *technicalTeamRepo) GetAll(ctx context.Context, pageSize int, page int) ([]domain.TechnicalTeam, error) {
	var technicalTeams []domain.TechnicalTeam
	err := r.db.Table("technical_teams").WithContext(ctx).Limit(pageSize).Offset(page - 1*pageSize).Find(&technicalTeams).Error
	if err != nil {
		return nil, err
	}
	return technicalTeams, nil
}
func (r *technicalTeamRepo) SetMember(ctx context.Context, teamId uuid.UUID, technicalTeamMember domain.TechnicalTeamMember) error {
	technicalTeam := mapper.TechnicalTeamMemberDomain2Storage(technicalTeamMember)
	return r.db.Table("technical_team_members").WithContext(ctx).Create(technicalTeam).Error
}
func (r *technicalTeamRepo) SetToTrip(ctx context.Context, teamId uuid.UUID, tripId uuid.UUID) error {
	return r.db.Table("technical_teams").WithContext(ctx).Where("id = ?", teamId).Update("trip_id", tripId).Error
}
