package utils

import (
	"github.com/shirou/gopsutil/process"
	"strings"
)

func IsProcessRunning(processName string) (bool, error) {
	processes, err := process.Processes()
	if err != nil {
		return false, err
	}

	for _, p := range processes {
		name, _ := p.Name()
		x := strings.TrimSuffix(name, ".exe")
		if strings.ToLower(x) == strings.ToLower(processName) {
			return true, nil
		}
	}

	return false, nil
}
