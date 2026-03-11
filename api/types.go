package api

type CPUInfo struct {
	Usage       float64 `json:"usage"`
	Temperature float64 `json:"temperature"`
}

type RAMInfo struct {
	Usage float64 `json:"usage"`
	Total uint64  `json:"total"`
}

type DiskInfo struct {
	Usage float64 `json:"usage"`
	Total uint64  `json:"total"`
}

type NetworkInfo struct {
	BytesSent     uint64 `json:"bytesSent"`
	BytesReceived uint64 `json:"bytesReceived"`
}
