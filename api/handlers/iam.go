package handlers

import (
	"gosecops/internal/cloud"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAMRequest struct {
	Profile string `json:"profile"`
}

func HandleIAMCheck(c *gin.Context) {
	var req IAMRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid profile"})
		return
	}
	results, err := cloud.AnalyzeIAMPolicies(req.Profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
