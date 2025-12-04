package main

import (
	"log"
	"net/http"

	"github.com/my-portfolio/backend/routes"
)

func main() {
	log.Println("Started execution of main...")

	http.HandleFunc("/health", routes.HealthCheck)
	http.HandleFunc("/dummy", routes.Dummy)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
