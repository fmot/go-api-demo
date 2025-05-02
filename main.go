package main

import (
	"example.com/go-api-demo/db"
	"example.com/go-api-demo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)


	server.Run(":8080")
}