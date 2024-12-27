package port

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/pb"
)

type GRPCBankClient interface {
	CreateUserWallet(userUUID string) (*pb.CreateWalletResponse, error)
	CreateFactor(rq *pb.CreateFactorRequest) (*pb.CreateFactorResponse, error)
}
