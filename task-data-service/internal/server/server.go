package server

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/config"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/handlers"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/logger"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/processors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/storage"
	api "github.com/dupreehkuda/TaskBingo/task-data-service/pkg/api"
)

// server provides single configuration out of all components
type server struct {
	handlers handlers.Handler
	config   *config.Config
	logger   *zap.Logger
}

// NewByConfig returns server instance with default config
func NewByConfig() *server {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	store := storage.New(cfg.DatabasePath, cfg.DatabasePass, log)

	proc := processors.New(store, log)

	handle := handlers.New(proc, log)

	return &server{
		handlers: handle,
		logger:   log,
		config:   cfg,
	}
}

// Run runs the service
func (a server) Run() {
	s := grpc.NewServer()
	serv := a.handlers
	api.RegisterTasksServer(s, serv)

	l, err := net.Listen("tcp", a.config.Address)
	if err != nil {
		a.logger.Fatal("server down")
	}

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

		s.GracefulStop()
		a.logger.Info("Server shut down", zap.String("port", a.config.Address))
		serverStopCtx()
	}()

	a.logger.Info("Server started", zap.String("port", a.config.Address))
	if err := s.Serve(l); err != nil {
		a.logger.Fatal("server down")
	}

	<-serverCtx.Done()
}
