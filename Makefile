
proto:
	protoc --go_out=internal/models --go_opt=paths=source_relative \
    --go-grpc_out=internal/models --go-grpc_opt=paths=source_relative api/api.proto

build:
	go build -o bin/client cmd/client/main.go
	go build -o bin/server cmd/server/main.go