package main

import (
	"flag"
	"log"

	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/server"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
