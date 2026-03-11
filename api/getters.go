package api

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/sensors"
)

func getCPUInfo() (CPUInfo, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return CPUInfo{}, err
	}

	temps, err := sensors.SensorsTemperatures()
	if err != nil {
		return CPUInfo{}, err
	}

	var temp float64
	for _, t := range temps {
		temp = t.Temperature
		break
	}

	return CPUInfo{
		Usage:       percent[0],
		Temperature: temp,
	}, nil
}

func getRAMInfo() (RAMInfo, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return RAMInfo{}, err
	}

	return RAMInfo{
		Usage: v.UsedPercent,
		Total: v.Total,
	}, nil
}

func getDiskInfo() (DiskInfo, error) {
	u, err := disk.Usage("/")
	if err != nil {
		return DiskInfo{}, err
	}

	return DiskInfo{
		Usage: u.UsedPercent,
		Total: u.Total,
	}, nil
}

func getNetworkInfo() (NetworkInfo, error) {
	io, err := net.IOCounters(true)
	if err != nil {
		return NetworkInfo{}, err
	}

	var bytesSent, bytesRecv uint64
	for _, n := range io {
		bytesSent += n.BytesSent
		bytesRecv += n.BytesRecv
	}

	return NetworkInfo{
		BytesSent:     bytesSent,
		BytesReceived: bytesRecv,
	}, nil
}
