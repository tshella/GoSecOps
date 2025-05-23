package email

import (
	"net"
	"strings"
)

type EmailAnalysis struct {
	SPF    string
	DKIM   string
	DMARC  string
	Domain string
}

func AnalyzeEmailSecurity(domain string) EmailAnalysis {
	txtRecords, _ := net.LookupTXT(domain)
	spf := "Not Found"
	for _, txt := range txtRecords {
		if strings.HasPrefix(txt, "v=spf1") {
			spf = txt
			break
		}
	}

	dmarc := "Not Found"
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	for _, txt := range dmarcRecords {
		if strings.HasPrefix(txt, "v=DMARC1") {
			dmarc = txt
			break
		}
	}

	return EmailAnalysis{
		SPF:    spf,
		DKIM:   "Not analyzed (future)", // For now
		DMARC:  dmarc,
		Domain: domain,
	}
}
