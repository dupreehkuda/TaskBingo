package taskRepository

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// TaskRepository is an interface for task data service
type TaskRepository interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(packIDs []string) (*[]models.TaskPack, error)
}

// taskRepository provides connection to task service
type taskRepository struct {
	conn   api.TasksClient
	logger *zap.Logger
}

// New returns an instance of userRepository
func New(address string, logger *zap.Logger) *taskRepository {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Error when connecting to user-service (not connected)")
	}

	client := api.NewTasksClient(conn)

	return &taskRepository{
		conn:   client,
		logger: logger,
	}
}
