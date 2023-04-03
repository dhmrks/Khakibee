package handler

import (
	"khakibee/site-api/store"
)

// Handler : hanldes http requests
type Handler struct {
	s	*store.Store
}

// New : creates new handler
func New(s *store.Store) *Handler {
	return &Handler{s: s}
}