package cmd

import (
	"fmt"

	"github.com/tshella/gosecops/internal/cloud"

	"github.com/spf13/cobra"
)

var dnsDomain string
var dnsWordlist []string

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Cloud DNS misconfiguration tools",
}

var dnsCheckCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Check subdomains for dangling CNAMEs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Scanning:", dnsDomain)
		results := cloud.ScanSubdomains(dnsDomain, dnsWordlist)
		for _, r := range results {
			fmt.Printf("Subdomain: %s\n", r.Subdomain)
			if r.Error != "" {
				fmt.Println("  ❌", r.Error)
				continue
			}
			if r.Dangling {
				fmt.Println("  ⚠️ Dangling CNAME:", r.CNAME)
			} else {
				fmt.Println("  ✅ Resolved:", r.CNAME, r.IP)
			}
		}
	},
}

func init() {
	dnsCmd.AddCommand(dnsCheckCmd)
	dnsCheckCmd.Flags().StringVar(&dnsDomain, "domain", "", "Target domain")
	dnsCheckCmd.Flags().StringSliceVar(&dnsWordlist, "subdomains", []string{"www", "api", "cdn", "dev"}, "Wordlist to scan")
	_ = dnsCheckCmd.MarkFlagRequired("domain")
	rootCmd.AddCommand(dnsCmd)
}
