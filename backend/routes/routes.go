package routes

import (
	"github.com/go-chi/chi/v5"
)

func Utils(router chi.Router) {
	router.HandleFunc("/dummy", dummy)
	router.HandleFunc("/health", healthCheck)
}

func Builder(router chi.Router) {
	router.Get("/builder", checkBuilder)
}
