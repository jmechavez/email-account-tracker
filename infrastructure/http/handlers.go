package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmechavez/email-account-tracker/internal/dto"
	"github.com/jmechavez/email-account-tracker/internal/ports/services"
)

type UserHandler struct {
	service services.UserService
}

// func (h UserHandler) Users(w http.ResponseWriter, r *http.Request) {
// 	users, err := h.service.Users()
// 	if err != nil {
// 		writeResponse(w, err.Code, err.AsMessage())
// 		return
// 	}
// 	writeResponse(w, http.StatusOK, users)
// }

func (h UserHandler) IdNo(w http.ResponseWriter, r *http.Request) {
	idNo := r.URL.Query().Get("id_no")
	if idNo == "" {
		idNo = r.URL.Query().Get("id")
	}

	if idNo != "" {
		user, err := h.service.IdNo(idNo)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
			return
		}
		writeResponse(w, http.StatusOK, user)
		return
	}
	{
		users, err := h.service.Users()
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
			return
		}
		writeResponse(w, http.StatusOK, users)
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idNo := vars["id_no"]
	var req dto.UserEmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	} else {
		req.IdNo = idNo
		user, err := h.service.CreateUser(req)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
			return
		}
		writeResponse(w, http.StatusCreated, user)
	}
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
