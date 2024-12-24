PROTOC_GEN_GO := protoc -I=./ --go_out=./ --go_opt=paths=source_relative
PROTOC_GEN_GRPC := --go-grpc_out=./ --go-grpc_opt=paths=source_relative

BANK_PROTO := ./api/pb/bank.proto

.PHONY: bank-proto
bank-proto:
	$(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC) $(BANK_PROTO)

.PHONY: proto-all
proto-all: bank-proto
