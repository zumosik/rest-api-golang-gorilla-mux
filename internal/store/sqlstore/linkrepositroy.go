package sqlstore

import "github.com/zumosik/rest-api-golang-gorilla-mux/internal/model"

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) FindByUser(u *model.User) ([]*model.Link, error) {
	query := "SELECT id, url, userid FROM links WHERE userid = $1"
	rows, err := r.store.db.Query(query, u.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*model.Link

	for rows.Next() {
		var link model.Link
		if err := rows.Scan(&link.ID, &link.URL, &link.UserID); err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	return links, nil
}

func (r *LinkRepository) Create(link *model.Link) error {
	if err := link.Validate(); err != nil {
		return err
	}

	query := "INSERT INTO links (url, userid) VALUES ($1, $2) RETURNING id"

	if err := r.store.db.QueryRow(query, link.URL, link.UserID).Scan(&link.ID); err != nil {

		return err
	}

	return nil

}
