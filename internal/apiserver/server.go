package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store"
)

const (
	sessionName        = "userid"
	CtxKeyUser  ctxKey = iota
	CtxKeyReqID
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errInvalidData              = errors.New("invalid data")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int

type server struct {
	router       *mux.Router
	store        store.Store
	sessionStore sessions.Store
	logger       *log.Logger
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		store:        store,
		router:       mux.NewRouter(),
		sessionStore: sessionStore,
		logger:       log.New(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/session", s.handleSessionCreate()).Methods("POST")

	// private routes
	private := s.router.PathPrefix("/link").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/new", s.handleLinkCreate()).Methods("POST")
	private.HandleFunc("/get", s.handleLinkGet()).Methods("GET")
}

// Routes

// Middlewares
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKeyReqID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger := log.WithFields(log.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(CtxKeyReqID),
		})
		logger.Debugf("Started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Debugf("Completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
	})
}

// Helpers
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
