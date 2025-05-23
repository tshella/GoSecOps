package api

import (
	"gosecops/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// 🔍 Port Scanner
		api.POST("/scan/port", handlers.HandlePortScan)

		// ✉️ Email Tools
		api.POST("/email/attack", handlers.HandleEmailAttack)
		api.POST("/email/analyze", handlers.HandleEmailAnalyze)

		// ☁️ Cloud Security Modules
		api.POST("/cloud/iam", handlers.HandleIAMCheck)
		api.POST("/cloud/s3", handlers.HandleS3Audit)
		api.POST("/cloud/dns", handlers.HandleDNSCloudScan)
	}

	return r
}
