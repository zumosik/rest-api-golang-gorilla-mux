package teststore

import (
	"strconv"

	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/model"
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/store"
)

// UserRepository ...
type LinkRepository struct {
	store *Store
	links map[int]*model.Link
}

// Create ...
func (r *LinkRepository) Create(link *model.Link) error {
	if err := link.Validate(); err != nil {
		return err
	}

	link.ID = len(r.links) + 1
	r.links[link.ID] = link

	return nil
}

func (r *LinkRepository) FindByUser(u *model.User) ([]*model.Link, error) {
	var links []*model.Link

	for _, l := range r.links {
		if l.UserID == strconv.Itoa(u.ID) {
			links = append(links, l)
		}
	}

	if len(links) < 1 {
		return nil, store.ErrRecordNotFound

	}

	return links, nil
}
