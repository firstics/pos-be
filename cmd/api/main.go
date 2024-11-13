package main

import (
	"log"

	_ "go.uber.org/automaxprocs"

	config "github.com/dataxai/inno-go-project-layout/pkg/config"
	di "github.com/dataxai/inno-go-project-layout/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeApp(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
