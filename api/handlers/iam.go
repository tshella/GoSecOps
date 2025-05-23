package handlers

import (
	"gosecops/internal/cloud"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAMRequest struct {
	Profile string `json:"profile"`
}

// HandleIAMCheck godoc
// @Summary Audit AWS IAM users and attached policies
// @Description Detects overly permissive IAM policies
// @Tags Cloud
// @Accept json
// @Produce json
// @Param input body IAMRequest true "IAM Profile Input"
// @Success 200 {array} cloud.IAMFinding
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cloud/iam [post]
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
