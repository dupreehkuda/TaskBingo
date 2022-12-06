package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a api) router() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/", a.handlers.Ping)
	})

	return r
}
