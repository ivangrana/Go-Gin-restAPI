package main

import (
	"Phinance/database"
	"Phinance/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DB = database.Init()
	r := gin.Default()
	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080") // Listen on port 8080
}
