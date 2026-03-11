package api

type CPUInfo struct {
	Usage       float64
	Temperature float64
}

type RAMInfo struct {
	Usage float64
	Total uint64
}

type DiskInfo struct {
	Usage float64
	Total uint64
}

type NetworkInfo struct {
	BytesSent     uint64
	BytesReceived uint64
}
