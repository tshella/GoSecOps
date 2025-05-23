package utils

import (
	"net"
	"strings"
)

func IsIP(target string) bool {
	return net.ParseIP(target) != nil
}

func NormalizeDomain(domain string) string {
	return strings.TrimSpace(strings.ToLower(domain))
}
