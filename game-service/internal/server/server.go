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
	i "github.com/dupreehkuda/TaskBingo/game-service/internal/interfaces"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/logger"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/middleware"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/processors"
	taskClient "github.com/dupreehkuda/TaskBingo/game-service/internal/taskClient"
	userClient "github.com/dupreehkuda/TaskBingo/game-service/internal/userClient"
)

type api struct {
	handlers   i.Handlers
	middleware i.Middleware
	config     *config.Config
	logger     *zap.Logger
}

func NewByConfig() *api {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	uc := userClient.New(cfg.UserServiceAddress, log)
	tc := taskClient.New(cfg.TaskServiceAddress, log)
	mv := middleware.New(log)

	logic := processors.New(uc, tc, log)

	handle := handlers.New(logic, log)

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
