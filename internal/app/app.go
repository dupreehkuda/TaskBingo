package app

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/config"
	"github.com/dupreehkuda/TaskBingo/internal/logger"
	httpserver "github.com/dupreehkuda/TaskBingo/internal/rpc/http"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/comments"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/friend"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/game"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/game_realtime"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/middleware"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/pack"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/solo"
	"github.com/dupreehkuda/TaskBingo/internal/rpc/http/user"

	friendAccept "github.com/dupreehkuda/TaskBingo/internal/usecases/friend_accept"
	friendAcceptStore "github.com/dupreehkuda/TaskBingo/internal/usecases/friend_accept/storage"
	friendDelete "github.com/dupreehkuda/TaskBingo/internal/usecases/friend_delete"
	friendDeleteStore "github.com/dupreehkuda/TaskBingo/internal/usecases/friend_delete/storage"
	friendRequest "github.com/dupreehkuda/TaskBingo/internal/usecases/friend_request"
	friendRequestStore "github.com/dupreehkuda/TaskBingo/internal/usecases/friend_request/storage"

	gameAccept "github.com/dupreehkuda/TaskBingo/internal/usecases/game_accept"
	gameAcceptStore "github.com/dupreehkuda/TaskBingo/internal/usecases/game_accept/storage"
	gameCreate "github.com/dupreehkuda/TaskBingo/internal/usecases/game_create"
	gameCreateStore "github.com/dupreehkuda/TaskBingo/internal/usecases/game_create/storage"
	gameDelete "github.com/dupreehkuda/TaskBingo/internal/usecases/game_delete"
	gameDeleteStore "github.com/dupreehkuda/TaskBingo/internal/usecases/game_delete/storage"
	gameGet "github.com/dupreehkuda/TaskBingo/internal/usecases/game_get"
	gameGetStore "github.com/dupreehkuda/TaskBingo/internal/usecases/game_get/storage"
	gameRealtime "github.com/dupreehkuda/TaskBingo/internal/usecases/game_realtime_action"
	gameRealtimeStore "github.com/dupreehkuda/TaskBingo/internal/usecases/game_realtime_action/storage"

	packDislike "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_dislike"
	packDislikeStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_dislike/storage"
	packGetMany "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_many"
	packGetManyStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_many/storage"
	packGetRated "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_rated"
	packGetRatedStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_rated/storage"
	packLike "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_like"
	packLikeStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_like/storage"
	packRate "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_rate"
	packRateStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_rate/storage"
	packSet "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_set"
	packSetStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_set/storage"
	packUnrate "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_unrate"
	packUnrateStore "github.com/dupreehkuda/TaskBingo/internal/usecases/pack_unrate/storage"

	userGetAll "github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_all"
	userGetAllStore "github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_all/storage"
	userGetData "github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_data"
	userGetDataStore "github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_data/storage"
	userLogin "github.com/dupreehkuda/TaskBingo/internal/usecases/user_login"
	userLoginStore "github.com/dupreehkuda/TaskBingo/internal/usecases/user_login/storage"
	userRegister "github.com/dupreehkuda/TaskBingo/internal/usecases/user_register"
	userRegisterStore "github.com/dupreehkuda/TaskBingo/internal/usecases/user_register/storage"
	userStats "github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats"
	userStatsStore "github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats/storage"

	soloFinish "github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_finish"
	soloFinishStore "github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_finish/storage"
	soloProgress "github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_progress"
	soloProgressStore "github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_progress/storage"
	soloStart "github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_start"
	soloStartStore "github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_start/storage"

	commentAdd "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_add"
	commentAddStore "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_add/storage"
	commentDelete "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_delete"
	commentDeleteStore "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_delete/storage"
	commentEdit "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_edit"
	commentEditStore "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_edit/storage"
	commentList "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_list"
	commentListStore "github.com/dupreehkuda/TaskBingo/internal/usecases/comment_list/storage"

	"github.com/dupreehkuda/TaskBingo/pkg/pgxtx"
)

type App struct {
	Cfg     *config.Config
	Logger  *zap.Logger
	Pool    *pgxpool.Pool
	Handler http.Handler
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log := logger.New(cfg.IsLocal())

	if err := MigrateUp(cfg.DatabaseDSN, cfg.MigrationsPath); err != nil {
		log.Error("migrate up failed", zap.Error(err))
		return nil, err
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	txm := pgxtx.New(pool)
	h := hasher{}
	ts := newTokenSigner(cfg.JWTSecret)

	// user
	register := userRegister.New(userRegisterStore.New(pool), h, ts, txm, log)
	login := userLogin.New(userLoginStore.New(pool), h, ts, log)
	getData := userGetData.New(userGetDataStore.New(pool, log), log)
	getAll := userGetAll.New(userGetAllStore.New(pool), log)
	stats := userStats.New(userStatsStore.New(pool, log), log)

	// pack
	pSet := packSet.New(packSetStore.New(pool), txm, log)
	pGetMany := packGetMany.New(packGetManyStore.New(pool), log)
	pGetRated := packGetRated.New(packGetRatedStore.New(pool), log)
	pLike := packLike.New(packLikeStore.New(pool), txm)
	pDislike := packDislike.New(packDislikeStore.New(pool), txm)
	pRate := packRate.New(packRateStore.New(pool), txm)
	pUnrate := packUnrate.New(packUnrateStore.New(pool), txm)

	// friend
	fReq := friendRequest.New(friendRequestStore.New(pool), txm)
	fAcc := friendAccept.New(friendAcceptStore.New(pool), txm)
	fDel := friendDelete.New(friendDeleteStore.New(pool), txm)

	// game
	gCreate := gameCreate.New(gameCreateStore.New(pool), txm, log)
	gGet := gameGet.New(gameGetStore.New(pool), log)
	gAcc := gameAccept.New(gameAcceptStore.New(pool))
	gDel := gameDelete.New(gameDeleteStore.New(pool))

	// realtime
	realtimeUC := gameRealtime.New(gameRealtimeStore.New(pool), txm, log)
	hub := game_realtime.NewHub(realtimeUC)
	wsHandler := game_realtime.NewHandler(hub, realtimeUC, cfg.AllowedWSOrigins, log)

	// solo
	sStart := soloStart.New(soloStartStore.New(pool), txm, log)
	sProgress := soloProgress.New(soloProgressStore.New(pool))
	sFinish := soloFinish.New(soloFinishStore.New(pool), txm, log)

	// comments
	cAdd := commentAdd.New(commentAddStore.New(pool), log)
	cList := commentList.New(commentListStore.New(pool), log)
	cEdit := commentEdit.New(commentEditStore.New(pool), log)
	cDel := commentDelete.New(commentDeleteStore.New(pool))

	routes := httpserver.Routes{
		UserRegister: user.NewRegisterHandler(register, cfg.CurrentDomain, log),
		UserLogin:    user.NewLoginHandler(login, cfg.CurrentDomain, log),
		UserGetData:  user.NewGetDataHandler(getData, log),
		UserGetAll:   user.NewGetAllHandler(getAll, log),
		UserStats:    user.NewStatsHandler(stats, log),

		PackSet:      pack.NewSetHandler(pSet, log),
		PackGetMany:  pack.NewGetManyHandler(pGetMany, log),
		PackGetRated: pack.NewGetRatedHandler(pGetRated, log),
		PackLike:     pack.NewLikeHandler(pLike, log),
		PackDislike:  pack.NewDislikeHandler(pDislike, log),
		PackRate:     pack.NewRateHandler(pRate, log),
		PackUnrate:   pack.NewUnrateHandler(pUnrate, log),

		FriendRequest: friend.NewRequestHandler(fReq, log),
		FriendAccept:  friend.NewAcceptHandler(fAcc, log),
		FriendDelete:  friend.NewDeleteHandler(fDel, log),

		GameCreate: game.NewCreateHandler(gCreate, log),
		GameGet:    game.NewGetHandler(gGet, log),
		GameAccept: game.NewAcceptHandler(gAcc, log),
		GameDelete: game.NewDeleteHandler(gDel, log),

		GameRealtime: wsHandler,

		SoloStart:    solo.NewStartHandler(sStart, log),
		SoloProgress: solo.NewProgressHandler(sProgress, log),
		SoloFinish:   solo.NewFinishHandler(sFinish, log),

		CommentAdd:    comments.NewAddHandler(cAdd, log),
		CommentList:   comments.NewListHandler(cList, log),
		CommentEdit:   comments.NewEditHandler(cEdit, log),
		CommentDelete: comments.NewDeleteHandler(cDel, log),
	}

	mw := middleware.New(log, cfg.JWTSecret)
	handler := httpserver.New(routes, mw, cfg.IsLocal())

	return &App{
		Cfg:     cfg,
		Logger:  log,
		Pool:    pool,
		Handler: handler,
	}, nil
}

func (a *App) Close() {
	if a.Pool != nil {
		a.Pool.Close()
	}
	if a.Logger != nil {
		_ = a.Logger.Sync()
	}
}
