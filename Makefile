
APP_NAME := "app"

.PHONY: build
build:
	go build -o $(APP_NAME) main.go
