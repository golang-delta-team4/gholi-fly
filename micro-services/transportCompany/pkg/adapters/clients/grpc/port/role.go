package port

import (
	rolePB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/role"
)

type GRPCRoleClient interface {
	CreateRole(rq *rolePB.GrantResourceAccessRequest) (*rolePB.GrantResourceAccessResponse, error)
}
