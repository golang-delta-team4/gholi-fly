package port

import (
	"user-service/pkg/adapters/clients/grpc/pb"
)

type GRPCBankClient interface {
	CreateUserWallet(userUUID string) (*pb.CreateWalletResponse, error)
}
