package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetTaskPacks handles getting several task packs operations
func (h *handlers) GetTaskPacks(w http.ResponseWriter, r *http.Request) {
	var req models.TaskPacksRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("Unable to read body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = easyjson.Unmarshal(body, &req); err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = UUIDCheck(req.PackIDs...); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.service.GetTaskPacks(r.Context(), req.PackIDs...)

	switch {
	case errors.Is(err, errs.ErrNoSuchPack):
		w.WriteHeader(http.StatusNoContent)
		return
	case err != nil:
		h.logger.Error("Unable to return task pack", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultJSON, err := easyjson.Marshal(resp)
	if err != nil {
		h.logger.Error("Error marshaling data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(resultJSON)
	if err != nil {
		h.logger.Error("Unable to write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
