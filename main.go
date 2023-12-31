package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/events-rest-api/db"
	"example.com/events-rest-api/routes"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.GET("/ping", pingHandler)

	routes.EventRouteHandlers(server)
	routes.UsersRoutes(server)

	server.Run(":8000")

}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
