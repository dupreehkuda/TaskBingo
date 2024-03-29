package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// router provides service main routing
func (a api) router() http.Handler {
	r := chi.NewRouter()

	if a.config.CurrentDomain == "localhost" {
		r.Use(a.middleware.RequestLogger)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Use(a.middleware.CheckCompression)
			r.Use(a.middleware.WriteCompressed)

			r.Group(func(r chi.Router) {
				r.Post("/register", a.handlers.RegisterUser)
				r.Post("/login", a.handlers.LoginUser)
			})

			r.Group(func(r chi.Router) {
				r.Use(a.middleware.CheckToken)

				r.Get("/getUserData", a.handlers.GetUserData)
				r.Post("/likePack", a.handlers.LikeTaskPack)
				r.Post("/dislikePack", a.handlers.DislikeTaskPack)
				r.Post("/ratePack", a.handlers.RateTaskPack)
				r.Post("/unratePack", a.handlers.UnrateTaskPack)
				r.Get("/getAllUsers", a.handlers.GetAllUsers)
				r.Post("/requestFriend", a.handlers.RequestFriend)
				r.Post("/acceptFriend", a.handlers.AcceptFriend)
				r.Post("/deleteFriend", a.handlers.DeleteFriend)
			})
		})

		r.Route("/task", func(r chi.Router) {
			r.Use(a.middleware.CheckToken)
			r.Use(a.middleware.CheckCompression)
			r.Use(a.middleware.WriteCompressed)

			r.Post("/getTaskPacks", a.handlers.GetTaskPacks)
			r.Post("/setTaskPack", a.handlers.SetTaskPack)
			r.Get("/getRatedPacks", a.handlers.GetRatedPacks)
		})

		r.Route("/game", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(a.middleware.AllowUpgradeHeaders)
				r.Get("/start", a.handlers.GameWSLaunch)
			})

			r.Group(func(r chi.Router) {
				r.Use(a.middleware.CheckToken)
				r.Use(a.middleware.CheckCompression)
				r.Use(a.middleware.WriteCompressed)

				r.Post("/get", a.handlers.GetGame)
				r.Post("/create", a.handlers.CreateGame)
				r.Patch("/accept", a.handlers.AcceptGame)
				r.Patch("/archive", nil)
				r.Patch("/update", nil)
				r.Delete("/delete", a.handlers.DeleteGame)
			})
		})
	})

	return r
}
