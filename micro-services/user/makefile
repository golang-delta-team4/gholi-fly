gen-proto:
	protoc -I=./api/pb -I=. --go_out=./api/pb --go_opt=paths=source_relative --go-grpc_out=./api/pb --go-grpc_opt=paths=source_relative ./api/pb/*.proto
