build:
	go build .

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protobuf/service.proto

run:
	go run main.go

grpc-ui:
	grpcui -plaintext localhost:50051

mod:
	go mod tidy