package main

import (
	"fmt"
	"net/http"
	"rest-go/internal/middleware"
	"rest-go/internal/user"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.JSONContentType)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!!!"))
	})

	user.SetupRoutes(r)

	fmt.Println("Starting server on :4545")
	if err := http.ListenAndServe(":4545", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
