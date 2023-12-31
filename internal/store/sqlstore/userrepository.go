package sqlstore

import (
	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {

		return err
	}

	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	if err := r.store.db.QueryRow(query, u.Email, u.Password).Scan(&u.ID); err != nil {

		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	query := "SELECT id, email, password FROM users WHERE email = $1"
	if err := r.store.db.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Password); err != nil {

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	query := "SELECT id, email, password FROM users WHERE id = $1"
	if err := r.store.db.QueryRow(query, id).Scan(&u.ID, &u.Email, &u.Password); err != nil {

		return nil, err
	}

	return u, nil
}
