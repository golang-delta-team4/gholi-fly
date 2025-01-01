package port

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/pb"
)

type GRPCBankClient interface {
	CreateFactor(rq *pb.CreateFactorRequest) (*pb.CreateFactorResponse, error)
	GetWallets(req *pb.GetWalletsRequest) (*pb.GetWalletsResponse, error)
	ApplyFactor(req *pb.ApplyFactorRequest) (*pb.ApplyFactorResponse, error)
	CancelFactor(req *pb.CancelFactorRequest) (*pb.CancelFactorResponse, error)
}
