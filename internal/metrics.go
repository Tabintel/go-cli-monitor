package internal

import(
	"github.com/shirou/gopsutil/cpu"
)

func GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}