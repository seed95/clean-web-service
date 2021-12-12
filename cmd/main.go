package main

import (
	"github.com/seed95/clean-web-service/internal/application"
	"github.com/seed95/clean-web-service/internal/config"
	"log"
)

var cfg = config.Config{}

var configPath = "./build/config/config.yaml"

func init() {
	if err := config.Parse(configPath, &cfg); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if err := application.Run(&cfg); err != nil {
		log.Fatalln(err)
	}
}
