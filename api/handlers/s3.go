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

func HandleS3Audit(c *gin.Context) {
	var req S3AuditRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Bucket == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid bucket name"})
		return
	}

	result := cloud.AuditS3Bucket(req.Profile, req.Bucket)
	c.JSON(http.StatusOK, result)
}
