package http

import (
	"encoding/json"
	"net/http"

	"github.com/jmechavez/email-account-tracker/internal/ports/services"
)

type UserHandler struct {
	service services.UserService
}

func (h UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.User()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, users)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow frontend
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
