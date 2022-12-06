package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/game-service/internal/interfaces"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

type handlers struct {
	userClient i.UserDataClient
	logger     *zap.Logger
}

// New creates new instance of handlers
func New(userClient i.UserDataClient, logger *zap.Logger) *handlers {
	return &handlers{
		userClient: userClient,
		logger:     logger,
	}
}

func (h handlers) Ping(w http.ResponseWriter, r *http.Request) {
	var data models.Request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := h.userClient.GetUser(data.UserID)
	if err != nil {
		h.logger.Error("Unable to call user microservice", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultJSON, err := json.Marshal(resp)
	if err != nil {
		h.logger.Error("Error marshaling data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	_, err = w.Write(resultJSON)
	if err != nil {
		h.logger.Error("Unable to write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Debug("everything good")
}
