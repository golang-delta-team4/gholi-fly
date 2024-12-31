package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
)

func TechnicalTeamDomain2Storage(domain domain.TechnicalTeam) *types.TechnicalTeam {
	return &types.TechnicalTeam{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		CompanyId:   domain.CompanyId,
		TripType:    domain.TripType,
	}
}

func TechnicalTeamStorage2Domain(storage types.TechnicalTeam) *domain.TechnicalTeam {
	return &domain.TechnicalTeam{
		Id:          storage.Id,
		Name:        storage.Name,
		Description: storage.Description,
		CompanyId:   storage.CompanyId,
		TripType:    storage.TripType,
		Members:     getIds(storage.Members),
	}
}

func TechnicalTeamMemberDomain2Storage(domain domain.TechnicalTeamMember) *types.TechnicalTeamMember {
	return &types.TechnicalTeamMember{
		UserId:          domain.UserId,
		Position:        domain.Position,
		TechnicalTeamId: domain.TeamId,
	}
}

func getIds(members []types.TechnicalTeamMember) []string {
	result := []string{}
	for _, item := range members {
		result = append(result, item.UserId.String())
	}

	return result
}
