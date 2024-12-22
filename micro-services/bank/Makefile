# Define protoc commands with options
PROTOC_GEN_GO := protoc -I=./ --go_out=./ --go_opt=paths=source_relative
PROTOC_GEN_GRPC := --go-grpc_out=./ --go-grpc_opt=paths=source_relative

# Proto file locations for each service
WALLET_PROTO := ./api/pb/wallet.proto
# Add additional proto files here
# USER_PROTO := ./api/pb/user.proto
# ORDER_PROTO := ./api/pb/order.proto

# Generate Wallet Protobuf
.PHONY: wallet-proto
wallet-proto:
	$(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC) $(WALLET_PROTO)

# Generate User Protobuf (example)
# .PHONY: user-proto
# user-proto:
# 	$(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC) $(USER_PROTO)

# Generate Order Protobuf (example)
# .PHONY: order-proto
# order-proto:
# 	$(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC) $(ORDER_PROTO)

# Generate all protobuf files
.PHONY: proto-all
proto-all: wallet-proto
# Add other proto targets here
#	proto-all: wallet-proto user-proto order-proto
