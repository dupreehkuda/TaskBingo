package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// LoginUser handles user login operations
func (h handlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var logCredit models.LoginCredentials

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&logCredit)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if logCredit.Username == "" && logCredit.Password == "" {
		h.logger.Info("Credentials empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.processor.LoginUser(logCredit.Username, logCredit.Password)

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
		Domain: "taskbingo.com",
		Path:   "/",
	})
}
