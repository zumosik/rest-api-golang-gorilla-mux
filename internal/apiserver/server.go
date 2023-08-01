package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store"
)

type server struct {
	router *mux.Router
	l      *log.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		store:  store,
		router: mux.NewRouter(),
		l:      log.New(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
