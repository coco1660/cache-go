package main

import (
	"github.com/coco1660/cache2go/config"
	"github.com/coco1660/cache2go/internal/app"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

func main() {
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(&cfg)
}
