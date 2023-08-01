package apiserver

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

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
