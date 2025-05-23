package cmd

import (
	"fmt"
	"gosecops/internal/cloud"

	"github.com/spf13/cobra"
)

var profile string

var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Cloud IAM analysis tools",
}

var iamCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Analyze AWS IAM policies for risks",
	Run: func(cmd *cobra.Command, args []string) {
		results, err := cloud.AnalyzeIAMPolicies(profile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, r := range results {
			fmt.Printf("User: %s | Policy: %s\n", r.User, r.PolicyName)
			for _, issue := range r.Issues {
				fmt.Println("  ⚠️", issue)
			}
		}
	},
}

func init() {
	iamCmd.AddCommand(iamCheckCmd)
	iamCheckCmd.Flags().StringVar(&profile, "profile", "default", "AWS CLI profile name")
	rootCmd.AddCommand(iamCmd)
}
