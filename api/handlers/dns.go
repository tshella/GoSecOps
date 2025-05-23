package handlers

import (
	"gosecops/internal/cloud"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DNSCloudRequest struct {
	Domain     string   `json:"domain"`
	Subdomains []string `json:"subdomains"`
}

func HandleDNSCloudScan(c *gin.Context) {
	var req DNSCloudRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if len(req.Subdomains) == 0 {
		req.Subdomains = []string{"www", "dev", "staging", "cdn", "api"}
	}

	results := cloud.ScanSubdomains(req.Domain, req.Subdomains)
	c.JSON(http.StatusOK, results)
}
