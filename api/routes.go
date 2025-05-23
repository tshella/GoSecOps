package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tshella/gosecops/api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/scan/port", handlers.HandlePortScan)
		api.POST("/email/attack", handlers.HandleEmailAttack)
		api.POST("/email/analyze", handlers.HandleEmailAnalyze)
		api.POST("/cloud/iam", handlers.HandleIAMCheck)
		api.POST("/cloud/s3", handlers.HandleS3Audit)
		api.POST("/cloud/dns", handlers.HandleDNSCloudScan)
	}

	return r
}
