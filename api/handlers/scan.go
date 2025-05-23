package handlers

import (
	"gosecops/internal/scanner"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScanRequest struct {
	Target string `json:"target"`
}

// HandlePortScan godoc
// @Summary Scan open TCP ports on a host
// @Description Performs a basic port scan (1-1024)
// @Tags Scanner
// @Accept json
// @Produce json
// @Param input body ScanRequest true "Target Host Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /scan/port [post]
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
