package port

import (
	bankPB "gholi-fly-hotel/pkg/adapters/clients/grpc/pb"
)

type GRPCBankClient interface {
	GetUserWallets(req *bankPB.GetWalletsRequest) (*bankPB.GetWalletsResponse, error)
	CreateFactor(req *bankPB.CreateFactorRequest) (*bankPB.CreateFactorResponse, error)
}
