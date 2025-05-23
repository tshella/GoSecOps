package main

import (
	"gosecops/api"
	_ "gosecops/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoSecOps API
// @version 1.0
// @description REST API for Penetration Testing & Cloud Security Tools
// @contact.name Manaka Anthony Raphasha
// @contact.email manaka@example.com
// @host localhost:8181
// @BasePath /api

func main() {
	r := api.SetupRouter()

	// Add Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8181")
}
