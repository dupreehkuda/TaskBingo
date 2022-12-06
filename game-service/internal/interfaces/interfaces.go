package interfaces

import (
	"net/http"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

type Handlers interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type UserDataClient interface {
	GetUser(userId string) (*models.Response, error)
}
