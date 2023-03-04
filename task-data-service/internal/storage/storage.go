package storage

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// Stored is interface for storage
type Stored interface {
	GetTaskPack(taskId string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
	CheckPackExistence(taskId string) (bool, error)
}

// storage provides a database connection
type storage struct {
	handle *rejson.Handler
	logger *zap.Logger
}

// New creates a new instance of database layer and migrates it
func New(path, pass string, logger *zap.Logger) *storage {
	// Wait until database initialize in container
	time.Sleep(1 * time.Second)
	rh := rejson.NewReJSONHandler()

	rdb := redis.NewClient(&redis.Options{
		Addr:     path,
		Password: pass, // no password set
		DB:       0,    // use default DB
	})

	rh.SetGoRedisClient(rdb)

	return &storage{
		handle: rh,
		logger: logger,
	}
}
