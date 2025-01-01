package port

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http/presenter"
	"github.com/google/uuid"
)

type HttpPathClient interface {
	GetPathDetail(pathID uuid.UUID) (*presenter.GetPathByIDResponse, error)
}
