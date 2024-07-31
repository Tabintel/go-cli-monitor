package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	//"github.com/tabintel/go-cli-monitor/internal"
)

var rootCmd = &cobra.Command{
	Use: "monitor",
	Short: "CLI tool for system monitoring",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Systsem monitoring CLI tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}