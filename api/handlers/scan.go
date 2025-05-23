package handlers

import (
	"gosecops/internal/scanner"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScanRequest struct {
	Target string `json:"target"`
}

func HandlePortScan(c *gin.Context) {
	var req ScanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	ports := scanner.PortScan(req.Target)
	c.JSON(http.StatusOK, gin.H{
		"target":     req.Target,
		"open_ports": ports,
	})
}
