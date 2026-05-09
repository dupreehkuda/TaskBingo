package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/comments"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/friend"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/game"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/game_realtime"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/middleware"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/pack"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/solo"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/user"
)

type Routes struct {
	UserRegister *user.RegisterHandler
	UserLogin    *user.LoginHandler
	UserGetData  *user.GetDataHandler
	UserGetAll   *user.GetAllHandler

	PackSet      *pack.SetHandler
	PackGetMany  *pack.GetManyHandler
	PackGetRated *pack.GetRatedHandler
	PackLike     *pack.LikeHandler
	PackDislike  *pack.DislikeHandler
	PackRate     *pack.RateHandler
	PackUnrate   *pack.UnrateHandler

	FriendRequest *friend.RequestHandler
	FriendAccept  *friend.AcceptHandler
	FriendDelete  *friend.DeleteHandler

	GameCreate *game.CreateHandler
	GameGet    *game.GetHandler
	GameAccept *game.AcceptHandler
	GameDelete *game.DeleteHandler

	GameRealtime *game_realtime.Handler

	SoloStart    *solo.StartHandler
	SoloProgress *solo.ProgressHandler
	SoloFinish   *solo.FinishHandler

	CommentAdd    *comments.AddHandler
	CommentList   *comments.ListHandler
	CommentEdit   *comments.EditHandler
	CommentDelete *comments.DeleteHandler
}

// New returns a chi router wired to the provided handlers and middleware,
// mirroring the legacy route tree.
func New(routes Routes, mw middleware.Middleware, isLocal bool) http.Handler {
	r := chi.NewRouter()

	if isLocal {
		r.Use(mw.RequestLogger)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Use(mw.CheckCompression)
			r.Use(mw.WriteCompressed)

			r.Group(func(r chi.Router) {
				r.Method(http.MethodPost, "/register", routes.UserRegister)
				r.Method(http.MethodPost, "/login", routes.UserLogin)
			})

			r.Group(func(r chi.Router) {
				r.Use(mw.CheckToken)
				r.Method(http.MethodGet, "/getUserData", routes.UserGetData)
				r.Method(http.MethodPost, "/likePack", routes.PackLike)
				r.Method(http.MethodPost, "/dislikePack", routes.PackDislike)
				r.Method(http.MethodPost, "/ratePack", routes.PackRate)
				r.Method(http.MethodPost, "/unratePack", routes.PackUnrate)
				r.Method(http.MethodGet, "/getAllUsers", routes.UserGetAll)
				r.Method(http.MethodPost, "/requestFriend", routes.FriendRequest)
				r.Method(http.MethodPost, "/acceptFriend", routes.FriendAccept)
				r.Method(http.MethodPost, "/deleteFriend", routes.FriendDelete)
			})
		})

		r.Route("/task", func(r chi.Router) {
			r.Use(mw.CheckToken)
			r.Use(mw.CheckCompression)
			r.Use(mw.WriteCompressed)

			r.Method(http.MethodPost, "/getTaskPacks", routes.PackGetMany)
			r.Method(http.MethodPost, "/setTaskPack", routes.PackSet)
			r.Method(http.MethodGet, "/getRatedPacks", routes.PackGetRated)
		})

		r.Route("/game", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(mw.AllowUpgradeHeaders)
				r.Method(http.MethodGet, "/start", routes.GameRealtime)
			})

			r.Group(func(r chi.Router) {
				r.Use(mw.CheckToken)
				r.Use(mw.CheckCompression)
				r.Use(mw.WriteCompressed)

				r.Method(http.MethodPost, "/get", routes.GameGet)
				r.Method(http.MethodPost, "/create", routes.GameCreate)
				r.Method(http.MethodPatch, "/accept", routes.GameAccept)
				r.Method(http.MethodDelete, "/delete", routes.GameDelete)

				r.Method(http.MethodPost, "/solo/start", routes.SoloStart)
				r.Method(http.MethodPut, "/solo/progress", routes.SoloProgress)
				r.Method(http.MethodPost, "/solo/finish", routes.SoloFinish)

				r.Method(http.MethodGet, "/{gameID}/comments", routes.CommentList)
				r.Method(http.MethodPost, "/{gameID}/comments", routes.CommentAdd)
				r.Method(http.MethodPatch, "/comments/{commentID}", routes.CommentEdit)
				r.Method(http.MethodDelete, "/comments/{commentID}", routes.CommentDelete)
			})
		})
	})

	return r
}
