package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmechavez/email-account-tracker/errors"
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

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {


	// Only allow PATCH method for updates
	if r.Method != http.MethodPatch {
		writeResponse(w, http.StatusMethodNotAllowed, errors.NewMethodNotAllowedError("Method not allowed"))
		return
	}

	// Extract IdNo from URL path
	vars := mux.Vars(r)
	idNo := vars["id_no"]
	if idNo == "" {
		writeResponse(w, http.StatusBadRequest, errors.NewBadRequestError("ID number is required in the URL"))
		return
	}

	// Parse the request body
	var request dto.UserUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errors.NewBadRequestError("Invalid request body"))
		return
	}

	// Assign the extracted IdNo to the request object
	request.IdNo = idNo

	// Call the service to update the user
	response, appError := h.service.UpdateUser(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError)
		return
	}

	// Return success response
	writeResponse(w, http.StatusOK, response)
}

func (h UserHandler) UpdateSurname(w http.ResponseWriter, r *http.Request) {
	// Only allow PATCH method for updates
	if r.Method != http.MethodPatch {
		writeResponse(w, http.StatusMethodNotAllowed, errors.NewMethodNotAllowedError("Method not allowed"))
		return
	}

	// Extract IdNo from URL path
	vars := mux.Vars(r)
	idNo := vars["id_no"]
	if idNo == "" {
		writeResponse(w, http.StatusBadRequest, errors.NewBadRequestError("ID number is required in the URL"))
		return
	}

	// Parse the request body
	var request dto.UserUpdateSurnameRequest
		err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errors.NewBadRequestError("Invalid request body"))
		return
	}

	// Assign the extracted IdNo to the request object
	request.IdNo = idNo

	// Call the service to update the user
	response, appError := h.service.UpdateSurname(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError)
		return
	}

	// Return success response
	writeResponse(w, http.StatusOK, response)
}

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

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idNo := vars["id_no"]
	var req dto.UserEmailDeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	} else {
		req.IdNo = idNo
		user, err := h.service.DeleteUser(req)
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
