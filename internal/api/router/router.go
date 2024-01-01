package router

import (
	"net/http"

	"example.com/internal/api/handler"
	"example.com/internal/api/repository"
	"example.com/internal/api/service"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	// Configure routes
	router := mux.NewRouter()

	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return router
}
