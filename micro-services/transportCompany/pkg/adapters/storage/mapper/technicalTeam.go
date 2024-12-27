package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
)

func TechnicalTeamDomain2Storage(domain domain.TechnicalTeam) types.TechnicalTeam {
	return types.TechnicalTeam{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		CompanyId:   domain.CompanyId,
		TripType:    domain.TripType,
	}
}

func TechnicalTeamMemberDomain2Storage(domain domain.TechnicalTeamMember) types.TechnicalTeamMember {
	return types.TechnicalTeamMember{
		UserId:          domain.UserId,
		Position:        domain.Position,
		TechnicalTeamId: domain.TeamId,
	}
}
