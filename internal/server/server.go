package server

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store"

	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
	l      *log.Logger
	router *mux.Router
	store  *store.Store
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

	if err := s.ConfigureStore(); err != nil {
		return err
	}

	s.l.Info("Starting server!")
	s.l.Debug("Bind address: ", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) ConfigureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *Server) ConfigureLogger() error {
	lvl, err := log.ParseLevel(s.config.LogLevel)
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
