package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/internal/api/router"
)

func main() {
	r := router.NewRouter()

  fmt.Println("Server started on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
