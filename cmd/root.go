package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosecops",
	Short: "GoSecOps CLI for security testing",
	Long:  `GoSecOps provides penetration testing and email security analysis via CLI.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(emailCmd)
}
