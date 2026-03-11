package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	cpuInfo, err := getCPUInfo()
	if err != nil {
		fmt.Fprintf(w, "Error getting CPU info: %v", err)
		return
	}

	ramInfo, err := getRAMInfo()
	if err != nil {
		fmt.Fprintf(w, "Error getting RAM info: %v", err)
		return
	}

	diskInfo, err := getDiskInfo()
	if err != nil {
		fmt.Fprintf(w, "Error getting disk info: %v", err)
		return
	}

	networkInfo, err := getNetworkInfo()
	if err != nil {
		fmt.Fprintf(w, "Error getting network info: %v", err)
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
