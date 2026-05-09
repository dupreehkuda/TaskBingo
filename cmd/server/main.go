package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/app"
	"github.com/dupreehkuda/TaskBingo/internal/config"
)

func main() {
	cfg := config.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, err := app.New(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer a.Close()

	srv := &http.Server{Addr: cfg.Address, Handler: a.Handler}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		a.Logger.Info("shutdown requested")
		shutdownCtx, c := context.WithTimeout(context.Background(), 30*time.Second)
		defer c()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			a.Logger.Error("shutdown error", zap.Error(err))
		}
		cancel()
	}()

	a.Logger.Info("server started", zap.String("addr", cfg.Address))
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.Logger.Fatal("listen failed", zap.Error(err))
	}
	<-ctx.Done()
}
