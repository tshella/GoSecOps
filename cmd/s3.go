package cmd

import (
	"fmt"
	"gosecops/internal/cloud"

	"github.com/spf13/cobra"
)

var s3Profile string
var s3Bucket string

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "S3 bucket audit tools",
}

var s3AuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Check if an S3 bucket is publicly exposed",
	Run: func(cmd *cobra.Command, args []string) {
		result := cloud.AuditS3Bucket(s3Profile, s3Bucket)
		fmt.Printf("Bucket: %s\n", result.Bucket)
		if result.Error != "" {
			fmt.Println("  ❌ Error:", result.Error)
			return
		}
		fmt.Printf("  Public Access Block Bypassed: %t\n", result.PublicAccess)
		fmt.Printf("  ACL Grants to AllUsers: %t\n", result.ACLGrantsExist)
	},
}

func init() {
	s3Cmd.AddCommand(s3AuditCmd)
	s3AuditCmd.Flags().StringVar(&s3Profile, "profile", "default", "AWS profile name")
	s3AuditCmd.Flags().StringVar(&s3Bucket, "bucket", "", "S3 bucket name")
	_ = s3AuditCmd.MarkFlagRequired("bucket")
	rootCmd.AddCommand(s3Cmd)
}
