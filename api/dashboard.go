package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const ERROR_STRING string = "error getting hardware info"

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	cpuInfo, err := getCPUInfo()
	if err != nil {
		log.Printf("error getting CPU info: %v\n", err)
		http.Error(w, ERROR_STRING, http.StatusInternalServerError)
		return
	}

	ramInfo, err := getRAMInfo()
	if err != nil {
		log.Printf("error getting RAM info: %v\n", err)
		http.Error(w, ERROR_STRING, http.StatusInternalServerError)
		return
	}

	diskInfo, err := getDiskInfo()
	if err != nil {
		log.Printf("error getting disk info: %v", err)
		http.Error(w, ERROR_STRING, http.StatusInternalServerError)
		return
	}

	networkInfo, err := getNetworkInfo()
	if err != nil {
		log.Printf("error getting disk info: %v\n", err)
		http.Error(w, ERROR_STRING, http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"cpu":     cpuInfo,
		"ram":     ramInfo,
		"disk":    diskInfo,
		"network": networkInfo,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
