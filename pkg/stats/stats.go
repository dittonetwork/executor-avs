package stats

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shirou/gopsutil/process"
)

// const bytesInMb = 1024 * 1024

func GetMemoryUsage() uint64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))             // go_memstats_alloc_bytes
	// fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc)) // go_memstats_alloc_bytes_total
	// fmt.Printf("\tSys = %v MiB", bToMb(memStats.Sys))               // go_memstats_sys_bytes
	// fmt.Printf("\tNumGC = %v\n", memStats.NumGC)             // go_gc_duration_seconds_count

	return memStats.Sys
}

// func bToMb(b uint64) uint64 {
// 	return b / bytesInMb
// }

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
