package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store"
)

type Store struct {
	db       *sql.DB
	userRepo *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepository{
		store: s,
	}

	return s.userRepo
}
