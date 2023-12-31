package router

import (
	"net/http"

	"example.com/internal/api/handler"
	"example.com/internal/api/repository"
	"example.com/internal/api/service"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	// Configurar rutas con Gorilla Mux
	router := mux.NewRouter()

	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	// router.HandleFunc("/users/{id}", handler.GetUser()).Methods("GET")
	// router.HandleFunc("/users", handler.AddUser()).Methods("POST")
	// router.HandleFunc("/users/{id}", handler.UpdateUser()).Methods("PUT")
	// router.HandleFunc("/users/{id}", handler.DeleteUser()).Methods("DELETE")

	return router
}
