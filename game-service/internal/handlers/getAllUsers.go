package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// GetAllUsers handles getting all users operation
func (h handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := h.processor.GetAllUsers()
	if err != nil {
		h.logger.Error("Error in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resultJSON, err := json.Marshal(resp)
	if err != nil {
		h.logger.Error("Error marshaling data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	_, err = w.Write(resultJSON)
	if err != nil {
		h.logger.Error("Unable to write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
