package stats

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shirou/gopsutil/process"
)

const bytesInMb = 1024 * 1024

func GetMemoryUsage() uint64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	// fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	// fmt.Printf("\tSys = %v MiB", bToMb(memStats.Sys))
	// fmt.Printf("\tNumGC = %v\n", memStats.NumGC)

	return bToMb(memStats.Sys)
}

func bToMb(b uint64) uint64 {
	return b / bytesInMb
}

func GetCPUUsage() (float64, error) {
	pid := int32(os.Getpid())
	proc, err := process.NewProcess(pid)
	if err != nil {
		return 0, fmt.Errorf("getting process: %w", err)
	}

	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		return 0, fmt.Errorf("getting cpu percent: %w", err)
	}

	return cpuPercent, nil
}
