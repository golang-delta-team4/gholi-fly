gen-proto:
	protoc -I=./api/pb -I=. --go_out=./api/pb --go_opt=paths=source_relative --go-grpc_out=./api/pb --go-grpc_opt=paths=source_relative ./api/pb/*.proto
gen-bank-proto:
	protoc -I=./pkg/adapters/clients/grpc/pb -I=. --go_out=./pkg/adapters/clients/grpc/pb --go_opt=paths=source_relative --go-grpc_out=./pkg/adapters/clients/grpc/pb --go-grpc_opt=paths=source_relative ./pkg/adapters/clients/grpc/pb/*.proto
