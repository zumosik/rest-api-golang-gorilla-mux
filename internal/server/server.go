package server

import (
	"os"

	log "github.com/sirupsen/logrus"

	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Server struct {
	config *Config
	l      *log.Logger
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		l:      log.New(),
	}
}

func (s *Server) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.l.Info("Starting server!")
	s.l.Error("ERROR")

	return nil
}

func (s *Server) ConfigureLogger() error {
	lvl, err := log.ParseLevel(s.config.Server.LogLevel)
	if err != nil {
		return err
	}
	s.l = &log.Logger{
		Out:   os.Stderr,
		Level: lvl,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}

	return nil
}
