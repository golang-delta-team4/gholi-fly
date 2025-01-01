package port

import (
	bankPB "gholi-fly-hotel/pkg/adapters/clients/grpc/pb"
)

type GRPCBankClient interface {
	GetUserWallets(req *bankPB.GetWalletsRequest) (*bankPB.GetWalletsResponse, error)
	CreateFactor(req *bankPB.CreateFactorRequest) (*bankPB.CreateFactorResponse, error)
	ApplyFactor(req *bankPB.ApplyFactorRequest) (*bankPB.ApplyFactorResponse, error)
	CancelFactor(req *bankPB.CancelFactorRequest) (*bankPB.CancelFactorResponse, error)
}
