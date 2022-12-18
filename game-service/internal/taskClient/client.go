package taskClient

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

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
