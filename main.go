package main

import (
	"log"
	"net/http"

	"github.com/karmek-k/rpi-dashboard/api"
)

func main() {
	http.HandleFunc("/", api.DashboardHandler)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
