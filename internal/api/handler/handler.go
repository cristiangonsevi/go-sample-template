package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/internal/api/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// getUsers recupera todos los usuarios
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers()

	if err != nil {
		handleError(w, err)
	}

	respondJSON(w, 200, users)

}

// getUser recupera un usuario por su ID
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
//
// }
//
// // addUser agrega un nuevo usuario
// func AddUser(w http.ResponseWriter, r *http.Request) {
// 	var user model.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
// }
//
// // updateUser actualiza un usuario existente por su ID
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
// }
//
// // deleteUser elimina un usuario por su ID
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
//
//
// }

// handleError maneja errores devolviendo una respuesta JSON
func handleError(w http.ResponseWriter, err error) {
	log.Println(err)
	respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error interno del servidor"})
}

// respondJSON envía una respuesta JSON con el código de estado especificado
func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
