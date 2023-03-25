package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/processors"
)

// Handlers is an interface for handlers
type Handlers interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)

	GetTaskPack(w http.ResponseWriter, r *http.Request)
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
}

// Handlers provides access to service
type handlers struct {
	processor processors.Processor
	logger    *zap.Logger
}

// New creates new instance of handlers
func New(processor processors.Processor, logger *zap.Logger) *handlers {
	return &handlers{
		processor: processor,
		logger:    logger,
	}
}

// UUIDCheck checks all needed ids to be sure of incoming data
func UUIDCheck(uuids ...string) error {
	for _, id := range uuids {
		_, err := uuid.Parse(id)
		if err != nil {
			return err
		}
	}

	return nil
}
