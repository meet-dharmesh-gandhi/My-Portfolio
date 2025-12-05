package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/meet-dharmesh-gandhi/my-portfolio/backend/routes"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// has the health check and dummmy api
	router.Route("/utils", routes.Utils)

	router.Route("/builder", routes.Builder)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
