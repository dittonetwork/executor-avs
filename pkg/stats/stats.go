package stats

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
)

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
