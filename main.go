package main

import (
	"github.com/tshella/gosecops/api"
	_ "github.com/tshella/gosecops/docs" // Needed for swag to include generated docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoSecOps API
// @version 1.0
// @description REST API for Penetration Testing & Cloud Security Tools
// @contact.name Manaka Anthony Raphasha
// @contact.email anthonyraphasha@gmail.com
// @host localhost:8181
// @BasePath /api
func main() {
	r := api.SetupRouter()

	// Serve Swagger UI at /swagger/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server on port 8181
	if err := r.Run(":8181"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
