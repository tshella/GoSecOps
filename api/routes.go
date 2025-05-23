package api

import (
	"gosecops/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/scan/port", handlers.HandlePortScan)
		// More routes will go here
	}
	return r
}
