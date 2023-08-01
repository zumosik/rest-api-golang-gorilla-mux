package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config   *Config
	db       *sql.DB
	userRepo *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {

		return err
	}

	if err := db.Ping(); err != nil { // check connection
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepository{
		store: s,
	}

	return s.userRepo
}
