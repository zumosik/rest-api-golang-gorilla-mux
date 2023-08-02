package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Link struct {
	ID     int    `json:"id"`
	URL    string `json:"URL"`
	UserID string `json:"-"`
}

func (link *Link) Validate() error {
	return validation.ValidateStruct(link, validation.Field(&link.URL, validation.Required, is.URL))
}
