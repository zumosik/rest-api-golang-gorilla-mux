package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	log.Info("Connected to db!")

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	log.Info("Starting server!")
	log.Debug("Address: ", config.BindAddr)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
