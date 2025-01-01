package port

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http/presenter"
)

type HttpVehicleClient interface {
	GetMatchedVehicle(matchMakerRequest *presenter.MatchMakerRequest) (*presenter.MatchMakerResponse, error)
}
