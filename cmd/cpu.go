package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tabintel/go-cli-monitor/internal"
)

var cpuCmd = &cobra.Command{
	Use: "cpu",
	Short: "Get CPU usage",
	Run: func(cmd *cobra.Command, args []string) {
		usage, err := internal.GetCPUUsage()
		if err != nil {
			fmt.Println("Error getting CPU usage:", err)
			return
		}
		fmt.Printf("CPU Usage: %.2f%%\n", usage)
	},
}