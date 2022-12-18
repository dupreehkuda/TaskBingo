package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// router provides service main routing
func (a api) router() http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Post("/register", a.handlers.RegisterUser)
				r.Post("/login", a.handlers.LoginUser)
			})

			r.Group(func(r chi.Router) {
				r.Use(a.middleware.CheckToken)

				r.Post("/getUserData", a.handlers.GetUserData)
			})
		})

		r.Route("/task", func(r chi.Router) {
			r.Use(a.middleware.CheckToken)

			r.Post("/getTaskPack", a.handlers.GetTaskPack)
			r.Post("/setTaskPack", a.handlers.SetTaskPack)
		})
	})

	return r
}
