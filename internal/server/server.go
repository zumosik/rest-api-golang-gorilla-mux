package server

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
	l      *log.Logger
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		l:      log.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.ConfigureRouter()

	s.l.Info("Starting server!")

	return http.ListenAndServe(s.config.Server.BindAddr, s.router)
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

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
