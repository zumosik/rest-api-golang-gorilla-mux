package store

import "github.com/zumosik/rest-api-golang-gorilla-mux/internal/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}

type LinkRepository interface {
	FindByUser(u *model.User) ([]*model.Link, error)
	Create(link *model.Link) error
}
