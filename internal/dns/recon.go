package dns

import (
	"fmt"
	"net"
)

func BruteSubdomains(domain string, wordlist []string) []string {
	found := []string{}
	for _, word := range wordlist {
		sub := fmt.Sprintf("%s.%s", word, domain)
		_, err := net.LookupHost(sub)
		if err == nil {
			found = append(found, sub)
		}
	}
	return found
}
