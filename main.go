/*
package main

import (
	
)

func main() {
	cmd.Execute()
}

*/

package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/spf13/cobra"
	"github.com/tabintel/go-cli-monitor/cmd"
    //"github.com/tabintel/go-cli-monitor/internal"
)

var cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "cpu_usage_percentage",
    Help: "Current CPU usage percentage",
})

func init() {
    prometheus.MustRegister(cpuUsage)
}

func startMetricsServer() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}

func main() {
    var rootCmd = &cobra.Command{
        Use:   "monitor",
        Short: "CLI tool for system monitoring",
    }

    var cpuCmd = &cobra.Command{
        Use:   "cpu",
        Short: "Get CPU usage",
        Run: func(cmd *cobra.Command, args []string) {
            usage, err := internal.GetCPUUsage()
            if err != nil {
                fmt.Println("Error getting CPU usage:", err)
                return
            }
            cpuUsage.Set(usage)
            fmt.Printf("CPU Usage: %.2f%%\n", usage)
        },
    }

    rootCmd.AddCommand(cpuCmd)
    go startMetricsServer()
    rootCmd.Execute()
}
