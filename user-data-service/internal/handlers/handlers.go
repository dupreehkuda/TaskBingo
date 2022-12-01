package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/internal/interfaces"
)

type handlers struct {
	storage i.Stored
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(storage i.Stored, logger *zap.Logger) *handlers {
	return &handlers{
		storage: storage,
		logger:  logger,
	}
}

type request struct {
	UserID string `json:"userID"`
}

func (h handlers) Ping(w http.ResponseWriter, r *http.Request) {
	var data request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := h.storage.Ping(data.UserID)
	if err != nil {
		h.logger.Error("Unable to call storage", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	_, err = w.Write(resp)
	if err != nil {
		h.logger.Error("Unable to write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Debug("everything good")
}
