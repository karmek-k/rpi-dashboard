package main

import (
	"log"
	"net/http"
	"os"

	"github.com/karmek-k/rpi-dashboard/api"
)

func main() {
	http.HandleFunc("/", api.DashboardHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":"+port, nil)
}
