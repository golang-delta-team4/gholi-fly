package port

import (
	bankPB "github.com/golang-delta-team4/gholi-fly-shared/pkg/protobuf/bank"
)

type GRPCBankClient interface {
	CreateUserWallet(userUUID string) (*bankPB.CreateWalletResponse, error)
}
