package handlers

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// LoginUser handles user login operations
func (h *handlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req models.LoginCredentials

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

	if req.Username == "" || req.Password == "" {
		h.logger.Info("Credentials empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.service.LoginUser(r.Context(), req.Username, req.Password)

	switch {
	case err == errs.ErrWrongCredentials:
		w.WriteHeader(http.StatusBadRequest)
		return
	case err != nil:
		h.logger.Error("Error in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  token,
		Secure: false,
		Domain: h.domain,
		Path:   "/",
	})
}
