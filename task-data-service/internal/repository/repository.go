package repository

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// Repository is interface for repository
type Repository interface {
	GetTaskPack(taskId string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
	CheckPackExistence(taskId string) (bool, error)
}

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
