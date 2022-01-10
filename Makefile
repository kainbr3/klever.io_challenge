build:
	go build .

build_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protobuf/service.proto

run:
	go run main.go
