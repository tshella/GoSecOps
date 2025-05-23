package cmd

import (
	"fmt"

	"github.com/tshella/gosecops/internal/email"

	"github.com/spf13/cobra"
)

var (
	from    string
	to      string
	subject string
	body    string
	domain  string
)

var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Email-related tools (attack, analyze)",
}

var emailAttackCmd = &cobra.Command{
	Use:   "attack",
	Short: "Send a spoofed email (for test use only!)",
	Run: func(cmd *cobra.Command, args []string) {
		err := email.SendSpoofedEmail(email.EmailAttackRequest{
			From:    from,
			To:      to,
			Subject: subject,
			Body:    body,
		})
		if err != nil {
			fmt.Println("❌ Failed to send email:", err)
			return
		}
		fmt.Printf("✅ Spoofed email sent from %s to %s\n", from, to)
	},
}

var emailAnalyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze SPF/DKIM/DMARC for a domain",
	Run: func(cmd *cobra.Command, args []string) {
		result := email.AnalyzeEmailSecurity(domain)
		fmt.Println("🔍 Analysis of:", domain)
		fmt.Println("SPF:   ", result.SPF)
		fmt.Println("DKIM:  ", result.DKIM)
		fmt.Println("DMARC: ", result.DMARC)
	},
}

func init() {
	// Attach subcommands to main "email" command
	emailCmd.AddCommand(emailAttackCmd)
	emailCmd.AddCommand(emailAnalyzeCmd)

	// Flags for attack
	emailAttackCmd.Flags().StringVar(&from, "from", "", "Spoofed sender email")
	emailAttackCmd.Flags().StringVar(&to, "to", "", "Recipient email")
	emailAttackCmd.Flags().StringVar(&subject, "subject", "Test Email", "Subject line")
	emailAttackCmd.Flags().StringVar(&body, "body", "This is a spoofed message.", "Email body")

	_ = emailAttackCmd.MarkFlagRequired("from")
	_ = emailAttackCmd.MarkFlagRequired("to")

	// Flags for analyze
	emailAnalyzeCmd.Flags().StringVar(&domain, "domain", "", "Target domain")
	_ = emailAnalyzeCmd.MarkFlagRequired("domain")
}
