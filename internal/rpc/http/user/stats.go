package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats/entity"
)

type StatsRunner interface {
	Run(ctx context.Context, userID string, days int) (*entity.Result, error)
}

type StatsHandler struct {
	uc     StatsRunner
	logger *zap.Logger
}

func NewStatsHandler(uc StatsRunner, logger *zap.Logger) *StatsHandler {
	return &StatsHandler{uc: uc, logger: logger}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)

	days := 7
	if raw := r.URL.Query().Get("days"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil {
			days = parsed
		}
	}

	data, err := h.uc.Run(r.Context(), userID, days)
	if err != nil {
		h.logger.Error("user_stats failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
