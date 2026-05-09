package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type Registrar interface {
	Run(ctx context.Context, c *models.RegisterCredentials) (string, error)
}

type RegisterHandler struct {
	uc     Registrar
	domain string
	logger *zap.Logger
}

func NewRegisterHandler(uc Registrar, domain string, logger *zap.Logger) *RegisterHandler {
	return &RegisterHandler{uc: uc, domain: domain, logger: logger}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var creds models.RegisterCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		h.logger.Error("decode register body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.uc.Run(r.Context(), &creds)
	if errors.Is(err, errs.ErrCredentialsInUse) {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		h.logger.Error("register failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	setAuthCookie(w, token, h.domain)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"username": creds.Username})
}

// setAuthCookie sets the auth JWT cookie. Used by both register and login handlers.
func setAuthCookie(w http.ResponseWriter, token, domain string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    token,
		Path:     "/",
		Domain:   domain,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
		Secure:   domain != "localhost",
		SameSite: http.SameSiteLaxMode,
	})
}
