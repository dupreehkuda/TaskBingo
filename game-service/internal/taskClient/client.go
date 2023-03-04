package taskClient

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// TaskDataClient is an interface for task data service
type TaskDataClient interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(packIDs []string) (*[]models.TaskPack, error)
}

// taskClient provides connection to task service
type taskClient struct {
	conn   api.TasksClient
	logger *zap.Logger
}

// New returns an instance of userClient
func New(address string, logger *zap.Logger) *taskClient {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Error when connecting to user-service (not connected)")
	}

	client := api.NewTasksClient(conn)

	return &taskClient{
		conn:   client,
		logger: logger,
	}
}
