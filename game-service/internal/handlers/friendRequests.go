package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// RequestFriend handles the operation of requesting friendship
func (h handlers) RequestFriend(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	if userID == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.FriendRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Person == "" {
		h.logger.Info("Request empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.processor.RequestFriend(userID, req.Person)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

// AcceptFriend handles the operation of accepting friendship
func (h handlers) AcceptFriend(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	if userID == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.FriendRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Person == "" {
		h.logger.Info("Request empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.processor.AcceptFriend(userID, req.Person)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

// DeleteFriend handles the operation of deleting friendship or canceling request
func (h handlers) DeleteFriend(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	if userID == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.FriendRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Person == "" {
		h.logger.Info("Request empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.processor.DeleteFriend(userID, req.Person)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}
