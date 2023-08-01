package main

import (
	"log"

	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/server"
)

func main() {
	configPath, err := server.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	config, err := server.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
