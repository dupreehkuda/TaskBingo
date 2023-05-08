package repository

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/zap"
)

// repository provides a database connection
type repository struct {
	handle *rejson.Handler
	logger *zap.Logger
}

// New creates a new instance of database layer and migrates it
func New(path, pass string, logger *zap.Logger) *repository {
	// Wait until database initialize in container
	time.Sleep(1 * time.Second)
	rh := rejson.NewReJSONHandler()

	rdb := redis.NewClient(&redis.Options{
		Addr:     path,
		Password: pass,
		DB:       0,
	})

	rh.SetGoRedisClient(rdb)

	return &repository{
		handle: rh,
		logger: logger,
	}
}
