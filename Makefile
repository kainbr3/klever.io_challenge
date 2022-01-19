build:
	go build .

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protobuf/service.proto

server:
	DATABASE_IP="localhost" go run main.go

client:
	go run src/command/client/client.go	

grpc-ui:
	grpcui -plaintext localhost:50051

mod:
	go mod tidy

web:
	dotnet watch -p src/frontend/kleverchallenge/kleverchallenge.csproj run

compose:
	docker compose build

up:
	docker compose up

list:
	docker ps