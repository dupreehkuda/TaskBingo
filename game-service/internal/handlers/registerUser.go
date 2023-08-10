package handlers

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// RegisterUser handles user registration operations
func (h *handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterCredentials

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

	if req.Username == "" || req.Password == "" || req.Email == "" {
		h.logger.Info("Credentials empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.service.RegisterUser(r.Context(), &req)

	switch {
	case err == errs.ErrCredentialsInUse:
		w.WriteHeader(http.StatusConflict)
		return
	case err != nil:
		h.logger.Error("Error in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")

	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  token,
		Secure: false,
		Domain: h.domain,
		Path:   "/",
	})
}
