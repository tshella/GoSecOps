package scanner

import (
	"fmt"
	"net"
	"time"
)

func PortScan(host string) []int {
	openPorts := []int{}
	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}
	return openPorts
}
