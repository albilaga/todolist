package main

import (
	"errors"
	"net/http"
)

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TodoRequest) Bind(r *http.Request) error {
	if t.Title == "" {
		return errors.New("missing title")
	}
	return nil
}
