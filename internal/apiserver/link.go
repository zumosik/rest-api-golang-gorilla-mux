package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zumosik/rest-api-golang-gorilla-mux/internal/model"
)

func (s *server) handleLinkCreate() http.HandlerFunc {
	type request struct {
		URL string `json:"url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, errInvalidData)
			return
		}

		u := r.Context().Value(CtxKeyUser).(*model.User)

		link := &model.Link{
			URL:    req.URL,
			UserID: strconv.Itoa(u.ID),
		}

		if err := s.store.Link().Create(link); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, link)

	}
}

func (s *server) handleLinkGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(CtxKeyUser).(*model.User)

		links, err := s.store.Link().FindByUser(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, links)
	}
}
