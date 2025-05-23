package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gosecops/internal/scanner"
)

var target string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run port scan",
	Run: func(cmd *cobra.Command, args []string) {
		results := scanner.PortScan(target)
		for _, r := range results {
			fmt.Println("Open port:", r)
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&target, "target", "t", "", "Target IP or hostname")
	scanCmd.MarkFlagRequired("target")
}
