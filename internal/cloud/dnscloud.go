package cloud

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type DNSCheckResult struct {
	Subdomain string `json:"subdomain"`
	CNAME     string `json:"cname,omitempty"`
	IP        string `json:"ip,omitempty"`
	Dangling  bool   `json:"dangling"`
	Error     string `json:"error,omitempty"`
}

func ScanSubdomains(domain string, wordlist []string) []DNSCheckResult {
	var results []DNSCheckResult

	for _, sub := range wordlist {
		full := fmt.Sprintf("%s.%s", sub, domain)
		result := DNSCheckResult{Subdomain: full}

		cname, err := net.LookupCNAME(full)
		if err == nil && cname != "" {
			result.CNAME = cname
			if strings.Contains(cname, "amazonaws.com") || strings.Contains(cname, "cloudfront.net") {
				// Potential for misconfig
				if isDangling(cname) {
					result.Dangling = true
				}
			}
		} else {
			ips, err := net.LookupIP(full)
			if err != nil {
				result.Error = "Unresolvable"
			} else if len(ips) > 0 {
				result.IP = ips[0].String()
			}
		}

		results = append(results, result)
	}

	return results
}

func isDangling(cname string) bool {
	client := net.Dialer{Timeout: 2 * time.Second}
	conn, err := client.Dial("tcp", cname+":80")
	if err != nil {
		return true
	}
	conn.Close()
	return false
}
