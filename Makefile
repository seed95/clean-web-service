
APP_NAME := "app"

.PHONY: build proto

build: proto
	go build -o $(APP_NAME) main.go

proto:
	protoc -I pb \
	--go_opt paths=source_relative \
	--go_out api \
	--go-grpc_opt paths=source_relative \
	--go-grpc_out api \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_out api \
	pb/**/**/*.proto
