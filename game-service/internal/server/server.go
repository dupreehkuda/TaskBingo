package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/config"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/handlers"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/logger"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/middleware"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/repository"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/service"
)

// Handlers is an interface for handlers
type Handlers interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)

	GetTaskPacks(w http.ResponseWriter, r *http.Request)
	SetTaskPack(w http.ResponseWriter, r *http.Request)
	LikeTaskPack(w http.ResponseWriter, r *http.Request)
	DislikeTaskPack(w http.ResponseWriter, r *http.Request)
	RateTaskPack(w http.ResponseWriter, r *http.Request)
	UnrateTaskPack(w http.ResponseWriter, r *http.Request)
	GetRatedPacks(w http.ResponseWriter, r *http.Request)

	GetAllUsers(w http.ResponseWriter, r *http.Request)
	RequestFriend(w http.ResponseWriter, r *http.Request)
	AcceptFriend(w http.ResponseWriter, r *http.Request)
	DeleteFriend(w http.ResponseWriter, r *http.Request)

	CreateGame(w http.ResponseWriter, r *http.Request)
	GetGame(w http.ResponseWriter, r *http.Request)
	AcceptGame(w http.ResponseWriter, r *http.Request)
	DeleteGame(w http.ResponseWriter, r *http.Request)

	GameWSLaunch(w http.ResponseWriter, r *http.Request)
}

// api provides single configuration out of all components
type api struct {
	handlers   Handlers
	middleware middleware.Middleware
	config     *config.Config
	logger     *zap.Logger
}

// NewByConfig returns server instance with default config
func NewByConfig() *api {
	cfg := config.New()
	log := logger.InitializeLogger(cfg.CurrentDomain == "localhost")

	ur := repository.New(cfg.UserServiceAddress, log)
	mv := middleware.New(log)

	logic := service.New(ur, log)

	handle := handlers.New(logic, cfg.CurrentDomain, log)

	return &api{
		handlers:   handle,
		middleware: mv,
		logger:     log,
		config:     cfg,
	}
}

// Run runs the service
func (a api) Run() {
	serv := &http.Server{Addr: a.config.Address, Handler: a.router()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				a.logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := serv.Shutdown(shutdownCtx)
		if err != nil {
			a.logger.Fatal("Error shutting down", zap.Error(err))
		}
		a.logger.Info("Server shut down", zap.String("port", a.config.Address))
		serverStopCtx()
	}()

	a.logger.Info("Server started", zap.String("port", a.config.Address))
	err := serv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		a.logger.Fatal("Cant start server", zap.Error(err))
	}

	<-serverCtx.Done()
}
