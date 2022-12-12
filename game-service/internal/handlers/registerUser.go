package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

func (h handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var regCredit models.RegisterCredentials

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&regCredit)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if regCredit.Login == "" && regCredit.Password == "" && regCredit.Email == "" {
		h.logger.Info("Credentials empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.processor.RegisterUser(regCredit.Login, regCredit.Email, regCredit.Password)

	switch {
	case err == errs.ErrCredentialsInUse:
		w.WriteHeader(http.StatusConflict)
		return
	case err != nil:
		h.logger.Error("Error in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: token,
	})
}
