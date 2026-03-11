package api

import (
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

	fmt.Fprintf(w, "CPU Usage: %.2f%%\n", cpuInfo.Usage)
	fmt.Fprintf(w, "CPU Temperature: %.2f°C\n", cpuInfo.Temperature)
	fmt.Fprintf(w, "RAM Usage: %.2f%%\n", ramInfo.Usage)
	fmt.Fprintf(w, "RAM Total: %.2f GB\n", float64(ramInfo.Total)/1024/1024/1024)
	fmt.Fprintf(w, "Disk Usage: %.2f%%\n", diskInfo.Usage)
	fmt.Fprintf(w, "Disk Total: %.2f GB\n", float64(diskInfo.Total)/1024/1024/1024)
	fmt.Fprintf(w, "Network Bytes Sent: %.2f MB\n", float64(networkInfo.BytesSent)/1024/1024)
	fmt.Fprintf(w, "Network Bytes Received: %.2f MB\n", float64(networkInfo.BytesReceived)/1024/1024)
}
