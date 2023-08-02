package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/apiserver"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		panic(err)
	}

	if err := configureLogger(config.LogLevel); err != nil {
		panic(err)
	}

	if err := apiserver.Start(config); err != nil {
		panic(err)
	}

}

func configureLogger(level string) error {

	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}

	log.SetLevel(lvl)

	log.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	return nil
}
