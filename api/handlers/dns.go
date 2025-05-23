package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tshella/gosecops/internal/cloud"
)

type DNSCloudRequest struct {
	Domain     string   `json:"domain"`
	Subdomains []string `json:"subdomains"`
}

// HandleDNSCloudScan godoc
// @Summary Scan cloud subdomains for misconfigurations
// @Description Detects dangling CNAMEs or unresolvable subdomains
// @Tags Cloud
// @Accept json
// @Produce json
// @Param input body DNSCloudRequest true "DNS Scan Input"
// @Success 200 {array} cloud.DNSCheckResult
// @Failure 400 {object} map[string]string
// @Router /cloud/dns [post]
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
