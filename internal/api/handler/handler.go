package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/internal/api/model"
	"example.com/internal/api/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers()

	if err != nil {
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, users)

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(w, err, ErrorResponse{Error: "Id must be an integer", Status: http.StatusBadRequest})
		return
	}

	user, err := h.UserService.GetUser(id)

	if err != nil {
		handleError(w, err, ErrorResponse{Error: "Record not found", Status: http.StatusNotFound})
		return
	}

	respondJSON(w, http.StatusOK, user)
}

// addUser create a new user
func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleError(w, err)
		return
	}
	err = h.UserService.AddUser(user)

	if err != nil {
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusCreated, "Record created successfully")
}

// updateUser
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(w, err, ErrorResponse{Error: "Id must be a number", Status: http.StatusBadRequest})
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		handleError(w, err)
		return
	}

	err = h.UserService.UpdateUser(id, user)

	respondJSON(w, http.StatusOK, "Record updated")

}

// deleteUser
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(w, err, ErrorResponse{Error: "Id must be a number", Status: http.StatusBadRequest})
		return
	}

	err = h.UserService.DeleteUser(id)

	if err != nil {
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, "Record deleted")

}

type ErrorResponse struct {
	Error  string
	Status int
}

func handleError(w http.ResponseWriter, err error, options ...ErrorResponse) {
	log.Println(err)

	status := http.StatusInternalServerError
	errMsg := "Internal Server Error"

	if len(options) > 0 {
		status = options[0].Status
		errMsg = options[0].Error
	}

	log.Println(errMsg)
	respondJSON(w, status, map[string]string{"error": errMsg})
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
