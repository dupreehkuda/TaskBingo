package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type Loginer interface {
	Run(ctx context.Context, username, password string) (string, error)
}

type LoginHandler struct {
	uc     Loginer
	domain string
	logger *zap.Logger
}

func NewLoginHandler(uc Loginer, domain string, logger *zap.Logger) *LoginHandler {
	return &LoginHandler{uc: uc, domain: domain, logger: logger}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var creds models.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := h.uc.Run(r.Context(), creds.Username, creds.Password)
	if errors.Is(err, errs.ErrWrongCredentials) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		h.logger.Error("login failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	setAuthCookie(w, token, h.domain)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"username": creds.Username})
}
