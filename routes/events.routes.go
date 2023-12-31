package routes

import (
	"example.com/events-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func EventRouteHandlers(server *gin.Engine) {

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", createEventHandler)
	authenticate.DELETE("/events/:id", deleteEventHandler)
	authenticate.PUT("/events/:id", updateEventHandler)

	server.GET("/events", getEventsHandler)
	server.GET("/events/:id", getEventHandler)
}
