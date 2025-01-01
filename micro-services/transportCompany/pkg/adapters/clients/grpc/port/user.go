package port

import (
	userPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/user"
)

type GRPCUserClient interface {
	CheckUserAuthorization(rq *userPB.UserAuthorizationRequest) (*userPB.UserAuthorizationResponse, error)
	GetBlockedUser(req *userPB.Empty) (*userPB.GetBlockedUsersResponse, error)
}
