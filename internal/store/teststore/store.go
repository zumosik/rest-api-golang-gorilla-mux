package teststore

import (
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/model"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store"
)

type Store struct {
	userRepository *UserRepository
	linkRepo       *LinkRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Link() store.LinkRepository {
	if s.linkRepo != nil {
		return s.linkRepo
	}

	s.linkRepo = &LinkRepository{
		store: s,
		links: make(map[int]*model.Link),
	}

	return s.linkRepo
}
