package port

import (
	userPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/user"
)

type GRPCUserClient interface {
	CreateRole(rq *userPB.UserAuthorizationRequest) (*userPB.UserAuthorizationResponse, error)
}
