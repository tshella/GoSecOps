package handlers

import (
	"gosecops/internal/cloud"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S3AuditRequest struct {
	Profile string `json:"profile"`
	Bucket  string `json:"bucket"`
}

// HandleS3Audit godoc
// @Summary Audit S3 bucket for public exposure
// @Description Checks ACLs and Public Access Block on given bucket
// @Tags Cloud
// @Accept json
// @Produce json
// @Param input body S3AuditRequest true "S3 Audit Request"
// @Success 200 {object} cloud.S3AuditResult
// @Failure 400 {object} map[string]string
// @Router /cloud/s3 [post]
func HandleS3Audit(c *gin.Context) {
	var req S3AuditRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Bucket == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid bucket name"})
		return
	}

	result := cloud.AuditS3Bucket(req.Profile, req.Bucket)
	c.JSON(http.StatusOK, result)
}
